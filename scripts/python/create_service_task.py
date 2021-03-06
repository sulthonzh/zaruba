from helper import cli
from helper.project import ServiceProject

@cli
def create_service_task(template_location: str, service_name: str, image_name: str, container_name: str, location: str, start_command: str='', ports: str='', envs: str='', dependencies: str='', runner_version: str=''):
    if location == '':
        raise 'Service location should be given'
    port_list = ports.split(',') if ports != '' else []
    env_list = envs.split(',') if envs != '' else []
    dependency_list = dependencies.split(',') if dependencies != '' else []
    dir_name = '.'
    service_project = ServiceProject()
    service_project.load_from_template(template_location)
    service_project.generate(dir_name=dir_name, service_name=service_name, image_name=image_name, container_name=container_name, location=location, start_command=start_command, port_list=port_list, env_list=env_list, dependency_list=dependency_list, runner_version=runner_version)


if __name__ == '__main__':
    create_service_task()