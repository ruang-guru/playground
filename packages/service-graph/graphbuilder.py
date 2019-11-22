import os
import yaml

class GraphBuilder:
    def __init__(self, mono_path):
        self.mono_path = mono_path

    def generate(self):
        nodes = {}
        projects = self.__list_projects()
        for project in projects:
            dep_protos = self.__get_dependencies(project)
            if dep_protos is not None:
                dep_projects = [proto.split('/')[0] for proto in dep_protos]
                nodes[project] = set(dep_projects)
            else:
                nodes[project] = set()

        return nodes

    def __list_projects(self):
        return [f for f in os.listdir(self.mono_path) if os.path.isdir(os.path.join(self.mono_path, f))]

    def __get_dependencies(self, project):
        project_path = os.path.join(self.mono_path, project)
        service_yaml_path = os.path.join(project_path, 'service.yaml')
        try:
            with open(service_yaml_path, 'r') as f:
                o = yaml.load(f,  Loader=yaml.BaseLoader)
                if 'dependencies' in o:
                    dependencies = o['dependencies']
                    return dependencies
                return None
        except FileNotFoundError:
            return None

