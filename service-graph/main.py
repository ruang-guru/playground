from graphbuilder import GraphBuilder
from dbretriever import DbRetriever

def render(g, schemas):
    from jinja2 import Environment, FileSystemLoader, select_autoescape
    from itertools import cycle, chain
    import json
    env = Environment(
        loader=FileSystemLoader('templates'),
        autoescape=select_autoescape(['html', 'xml'])
    )
    template = env.get_template('viz.html')
    nodes = [{'id': node, 'label': node, 'title': node} for node in g.keys()]
    edges =  [{'from': o[0], 'to': o[1]} for o in chain(*[list(zip(cycle([k]), v)) for (k, v) in g.items()])]
    print(template.render(nodes=nodes, edges=edges, schemas=schemas))

if __name__ == "__main__":
    builder = GraphBuilder('/Users/levifikri/go/src/gitlab.com/ruangguru/source')
    g = builder.generate()
    schemas = {}
    for project in g:
        db = DbRetriever(project)
        schemas[project] = db.retrieve_schema()
    render(g, schemas)
