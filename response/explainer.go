package response

import (
	"fmt"
	"sort"
	"strings"

	"github.com/state-alchemists/zaruba/config"
	"github.com/state-alchemists/zaruba/monitor"
	"github.com/state-alchemists/zaruba/str"
)

type Explainer struct {
	logger  monitor.Logger
	d       *monitor.Decoration
	project *config.Project
}

func NewExplainer(logger monitor.Logger, decoration *monitor.Decoration, project *config.Project) *Explainer {
	return &Explainer{
		logger:  logger,
		d:       decoration,
		project: project,
	}
}

func (e *Explainer) listToStr(list []string) string {
	if len(list) == 0 {
		return ""
	}
	separator := fmt.Sprintf("%s,%s ", e.d.Blue, e.d.Normal)
	return fmt.Sprintf("%s[ %s%s%s ]%s", e.d.Blue, e.d.Normal, strings.Join(list, separator), e.d.Blue, e.d.Normal)
}

func (e *Explainer) getStrOrBlank(str string) string {
	if str == "" {
		return fmt.Sprintf("%sBlank%s", e.d.Blue, e.d.Normal)
	}
	return str
}

func (e *Explainer) getFieldKeys(list []string) (keys []string) {
	keys = []string{}
	maxLength := 0
	for _, key := range list {
		if len(key) > maxLength {
			maxLength = len(key)
		}
	}
	for _, key := range list {
		fieldKey := key + strings.Repeat(" ", maxLength-len(key))
		keys = append(keys, fieldKey)
	}
	return keys
}

func (e *Explainer) Explain(taskName string) {
	task := e.project.Tasks[taskName]
	indentation := strings.Repeat(" ", 21)
	parentTasks := task.Extends
	if len(parentTasks) == 0 && task.Extend != "" {
		parentTasks = []string{task.Extend}
	}
	e.printField("TASK NAME   ", taskName, indentation)
	e.printField("LOCATION    ", task.GetFileLocation(), indentation)
	e.printField("DESCRIPTION ", task.Description, indentation)
	e.printField("DEPENDENCIES", e.listToStr(task.Dependencies), indentation)
	e.printField("PARENT TASKS", e.listToStr(parentTasks), indentation)
	e.printField("INPUTS      ", e.getInputString(task), indentation)
	e.printField("CONFIG      ", e.getConfigString(task), indentation)
	e.printField("LCONFIG     ", e.getLConfigString(task), indentation)
	e.printField("ENVIRONMENTS", e.getEnvString(task), indentation)
}

func (e *Explainer) getInputString(task *config.Task) (inputString string) {
	inputNames := task.Inputs
	inputCount := len(inputNames)
	if inputCount == 0 {
		return ""
	}
	for _, inputName := range inputNames {
		input := e.project.Inputs[inputName]
		rawInputLines := []string{
			inputName,
			e.getSubFieldString("DESCRIPTION", input.Description),
			e.getSubFieldString("PROMPT     ", input.Prompt),
			e.getSubFieldString("OPTIONS    ", e.listToStr(input.Options)),
			e.getSubFieldString("DEFAULT    ", input.DefaultValue),
			e.getSubFieldString("VALIDATION ", input.Validation),
		}
		inputLines := []string{}
		for _, inputLine := range rawInputLines {
			if inputLine != "" {
				inputLines = append(inputLines, inputLine)
			}
		}
		inputString += strings.Trim(strings.Join(inputLines, "\n"), "\n") + "\n"
	}
	return inputString
}

func (e *Explainer) getEnvString(task *config.Task) (envString string) {
	keys := task.GetEnvKeys()
	sort.Strings(keys)
	for _, key := range keys {
		env, _ := task.GetEnvObject(key)
		rawEnvLines := []string{
			key,
			e.getSubFieldString("FROM   ", env.From),
			e.getSubFieldString("DEFAULT", env.Default),
		}
		envLines := []string{}
		for _, envLine := range rawEnvLines {
			if envLine != "" {
				envLines = append(envLines, envLine)
			}
		}
		envString += strings.Trim(strings.Join(envLines, "\n"), "\n") + "\n"
	}
	return envString
}

func (e *Explainer) getConfigString(task *config.Task) (configStr string) {
	keys := task.GetConfigKeys()
	sort.Strings(keys)
	fieldKeys := e.getFieldKeys(keys)
	lines := []string{}
	for index, key := range keys {
		fieldKey := fieldKeys[index]
		val, _ := task.GetConfigPattern(key)
		fieldVal := e.getStrOrBlank(val)
		lines = append(lines, e.getSubFieldString(fieldKey, fieldVal))
	}
	return strings.Join(lines, "\n")
}

func (e *Explainer) getLConfigString(task *config.Task) (configStr string) {
	keys := task.GetLConfigKeys()
	sort.Strings(keys)
	fieldKeys := e.getFieldKeys(keys)
	lines := []string{}
	for index, key := range keys {
		vals, _ := task.GetLConfigPatterns(key)
		fieldKey := fieldKeys[index]
		fieldVal := e.listToStr(vals)
		if fieldVal == "" {
			fieldVal = fmt.Sprintf("%s[]%s", e.d.Blue, e.d.Normal)
		}
		lines = append(lines, e.getSubFieldString(fieldKey, fieldVal))
	}
	return strings.Join(lines, "\n")
}

func (e *Explainer) getSubFieldString(subFieldName string, subFieldValue string) (subFieldStr string) {
	subFieldIndentation := strings.Repeat(" ", 2)
	subFieldValue = strings.Trim(subFieldValue, "\n")
	if subFieldValue == "" {
		return ""
	}
	subFieldValueIndentation := subFieldIndentation + strings.Repeat(" ", len(subFieldName)+1)
	subFieldLines := strings.Split(strings.Trim(subFieldValue, "\n"), "\n")
	subFieldValueStr := strings.Join(subFieldLines, "\n  "+subFieldValueIndentation)
	return fmt.Sprintf("%s%s%s :%s %s", subFieldIndentation, e.d.Yellow, subFieldName, e.d.Normal, subFieldValueStr)
}

func (e *Explainer) printField(fieldName string, value string, indentation string) {
	trimmedValue := strings.TrimRight(value, "\n ")
	if trimmedValue == "" {
		return
	}
	valueLines := strings.Split(trimmedValue, "\n")
	indentedValue := strings.Join(valueLines, "\n"+indentation)
	e.logger.DPrintf("%s%s :%s %s\n", e.d.Yellow, fieldName, e.d.Normal, indentedValue)
}

func (e *Explainer) GetCommand(taskNames []string) (command string) {
	command = fmt.Sprintf("zaruba please %s", strings.Join(taskNames, " "))
	inputArgs := []string{}
	for _, taskName := range taskNames {
		task := e.project.Tasks[taskName]
		for _, inputName := range task.Inputs {
			inputValue := ""
			if input := e.project.Inputs[inputName]; input.Secret {
				inputValue = "[HIDDEN_VALUE]"
			} else {
				inputValue = e.project.GetValue(inputName)
			}
			inputArgs = append(inputArgs, fmt.Sprintf("%s=%s", inputName, str.DoubleQuoteShellValue(inputValue)))
		}
	}
	if len(inputArgs) != 0 {
		command = fmt.Sprintf("%s %s", command, strings.Join(inputArgs, " "))
	}
	return command
}
