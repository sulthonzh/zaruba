from typing import Mapping, Any, List, Tuple
import os
from project.structures import Task, ProjectDict
from ruamel.yaml import YAML


def write_task(file_name: str, task_name: str, task: Task):
    project_dict: ProjectDict = {}
    try:
        project_dict = get_dict_from_file(file_name)
    except:
        pass
    if 'tasks' not in project_dict:
        project_dict['tasks'] = {}
    if task_name in project_dict.tasks:
        return False
    project_dict['tasks'][task_name] = task.as_dict()
    save_dict_to_file(file_name, project_dict)


def add_to_main_include(file_name: str) -> bool:
    main_file_name = get_main_file_name()
    main_project_dict = get_dict_from_file(main_file_name)
    if 'includes' not in main_project_dict:
        main_project_dict['includes'] = []
    if file_name in main_project_dict['includes']:
        return False
    main_project_dict['includes'].append(file_name)
    save_dict_to_file(main_file_name, main_project_dict)
    return True


def add_to_main_task(task_name: str) -> bool:
    main_file_name = get_main_file_name()
    main_project_dict = get_dict_from_file(main_file_name)
    if 'tasks' not in main_project_dict:
        main_project_dict['tasks'] = {}
    task = Task(main_project_dict['tasks']['run']) \
        if 'run' in main_project_dict['tasks'] \
        else Task().set_icon('🚅').set_description('Run everything at once')
    task.add_dependency(task_name)
    main_project_dict['tasks']['run'] = task.as_dict()
    save_dict_to_file(main_file_name, main_project_dict)
    return True


def get_main_file_name() -> str:
    main_file_name = 'main.zaruba.yaml'
    return main_file_name


def create_dir(dirname: str):
    if not os.path.exists(dirname):
        os.makedirs(dirname)


def get_dict_from_file(file_name: str) -> Mapping[str, Any]:
    yaml=YAML()
    f = open(file_name, 'r')
    template_obj = yaml.load(f)
    f.close()
    return template_obj


def save_dict_to_file(file_name: str, dictionary: Mapping[str, Any]):
    yaml=YAML()
    f = open(file_name, 'w')
    yaml.dump(dictionary, f)
    f.close()


def write_task_env(file_name: str, task: Task):
    f = open(file_name, 'a')
    f.write('\n')
    for _, env in task.get_all_env().items():
        envvar = env.get_from()
        if envvar == '':
            continue
        value = env.get_default()
        f.write('{}={}\n'.format(envvar, value))
    f.close()


def get_default_template() -> str:
    return os.path.join(
        os.path.dirname(os.path.dirname(os.path.dirname(os.path.dirname(__file__)))),
        'templates'
    )


def get_template(template_path_list: List[str], location: str, default_location: str, location_is_dir=False) -> Tuple[str, bool]:
    real_template_path_list = [os.path.join('.', 'templates')]
    real_template_path_list.extend(template_path_list)
    default_template_location = get_default_template()
    real_template_path_list.append(default_template_location)
    for template_path in real_template_path_list:
        template_location = os.path.join(template_path, location)
        if (location_is_dir and os.path.isdir(template_location)) or os.path.isfile(template_location):
            return template_location, False
    return os.path.join(default_template_location, default_location), True


def get_service_task_template(template_path_list: List[str], service_type: str) -> Tuple[str, bool]:
    service_task_template_dir = 'service-task'
    location = os.path.join(service_task_template_dir, '{}.zaruba.yaml'.format(service_type))
    default_location = os.path.join(service_task_template_dir, 'default.zaruba.yaml')
    return get_template(template_path_list, location, default_location, location_is_dir=False)


def get_docker_task_template(template_path_list: List[str], image: str) -> Tuple[str, bool]:
    docker_task_template_dir = 'docker-task'
    location = os.path.join(docker_task_template_dir, '{}.zaruba.yaml'.format(image))
    default_location = os.path.join(docker_task_template_dir, 'default.zaruba.yaml')
    return get_template(template_path_list, location, default_location, location_is_dir=False)


def get_service_template(template_path_list: List[str], service_type: str) -> Tuple[str, bool]:
    service_template_dir = 'service'
    location = os.path.join(service_template_dir, service_type)
    default_location = os.path.join(service_template_dir, 'fastapi')
    return get_template(template_path_list, location, default_location, location_is_dir=True)
