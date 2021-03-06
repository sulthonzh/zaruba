# serveHttp
```
  TASK NAME     : serveHttp
  LOCATION      : /home/gofrendi/zaruba/scripts/core.zaruba.yaml
  DESCRIPTION   : Run static web server from your working directory.
  TASK TYPE     : Service Task
  PARENT TASKS  : [ core.startService ]
  START         : - python
                  - -m
                  - http.server
                  - {{ index (.Split (.Trim (.GetConfig "ports") "\n ") "\n") 0 }}
  CHECK         : - {{ .GetConfig "cmd" }}
                  - {{ .GetConfig "cmdArg" }}
                  - {{- $d := .Decoration -}}
                    {{ if .IsFalse (.GetConfig "runLocally") -}}
                      echo 🎉🎉🎉
                      echo "📜 {{ $d.Bold }}{{ $d.Yellow }}Task '{{ .Name }}' is ready{{ $d.Normal }}"
                      exit 0
                    {{ end -}}
                    {{ .Trim (.GetConfig "_setup") "\n " }}
                    {{ .Trim (.GetConfig "setup") "\n " }}
                    {{ .Trim (.GetConfig "beforeCheck") "\n " }}
                    {{ .Trim (.GetConfig "_check") "\n " }}
                    {{ .Trim (.GetConfig "check") "\n " }}
                    {{ .Trim (.GetConfig "afterCheck") "\n " }}
                    {{ .Trim (.GetConfig "finish") "\n " }}
                    echo 🎉🎉🎉
                    echo "📜 {{ $d.Bold }}{{ $d.Yellow }}Task '{{ .Name }}' is ready{{ $d.Normal }}"
  INPUTS        : server.httpPort
                    DESCRIPTION : HTTP port to serve static files
                    PROMPT      : HTTP port
                    OPTIONS     : [ 8080, 8000, 3000, 5000 ]
                    DEFAULT     : 8080
                    VALIDATION  : ^[0-9]+$
  CONFIG        : _setup                 : set -e
                                           alias zaruba=${ZARUBA_HOME}/zaruba
                                           {{ .Trim (.GetConfig "includeBootstrapScript") "\n" }}
                                           {{ .Trim (.GetConfig "includeUtilScript") "\n" }}
                  _start                 : Blank
                  afterCheck             : Blank
                  afterStart             : Blank
                  beforeCheck            : Blank
                  beforeStart            : Blank
                  check                  : {{- $d := .Decoration -}}
                                           {{ range $index, $port := .Split (.Trim (.GetConfig "ports") "\n ") "\n" -}}
                                             {{ if ne $port "" -}}
                                               echo "📜 {{ $d.Bold }}{{ $d.Yellow }}Waiting for port '{{ $port }}'{{ $d.Normal }}"
                                               wait_port "localhost" {{ $port }}
                                               echo "📜 {{ $d.Bold }}{{ $d.Yellow }}Port '{{ $port }}' is ready{{ $d.Normal }}"
                                             {{ end -}}
                                           {{ end -}}
                  cmd                    : {{ if .GetValue "defaultShell" }}{{ .GetValue "defaultShell" }}{{ else }}bash{{ end }}
                  cmdArg                 : -c
                  finish                 : Blank
                  includeBootstrapScript : if [ -f "${HOME}/.profile" ]
                                           then
                                               . "${HOME}/.profile"
                                           fi
                                           if [ -f "${HOME}/.bashrc" ]
                                           then
                                               . "${HOME}/.bashrc"
                                           fi
                                           BOOTSTRAP_SCRIPT="${ZARUBA_HOME}/scripts/bootstrap.sh"
                                           . "${BOOTSTRAP_SCRIPT}"
                  includeUtilScript      : . "${ZARUBA_HOME}/scripts/util.sh"
                  playBellScript         : echo $'\a'
                  ports                  : {{ .GetValue "server.httpPort" }}
                  runLocally             : true
                  setup                  : Blank
                  start                  : Blank
  ENVIRONMENTS  : PYTHONUNBUFFERED
                    FROM    : PYTHONUNBUFFERED
                    DEFAULT : 1
```
