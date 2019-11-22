import subprocess
import json
import psycopg2

class DbRetriever:
    def __init__(self, project):
        self.project = project
        self.db_host = None
        self.db_name = None
        self.db_pass = None
        self.db_username = None
        self.db_port = None

    def retrieve_schema(self):
        self.retrive_env()

        # print(self.db_host, self.db_name, self.db_username, self.db_port)
        res = []
        if self.db_host is not None:
            try:
                conn = psycopg2.connect(
                    database=self.db_name,
                    user=self.db_username,
                    sslmode='require',
                    port=self.db_port,
                    host=self.db_host,
                    password=self.db_pass
                )
                with conn.cursor() as cur:
                    cur.execute("SELECT name, database_name FROM crdb_internal.tables WHERE database_name = %s", (self.db_name, ))
                    tables = [item[0] for item in cur.fetchall()]
                    for table in tables:
                        cur.execute("SELECT * FROM public.{} LIMIT 5".format(table))
                        rows = [[str(col) for col in row] for row in cur.fetchall()]
                        cur.execute("SELECT column_name, column_type FROM crdb_internal.table_columns WHERE descriptor_name = %s", (table, ))
                        columns = [{"name": item[0], "type": item[1]} for item in cur.fetchall()]
                        res.append({"table": table, "columns": columns, "rows": rows})
                        # print(table, columns)
            except psycopg2.OperationalError:
                return []

        return res

    def retrive_env(self):
        res = subprocess.run(['rogu', 'config', 'preview', '-e', 'staging', '-s',
            self.project, '-c', 'id'], capture_output = True)
        out = res.stdout.decode('utf-8')
        if out:
            o = json.loads(out)
            if o.get('id'):
                self.db_host = o['id'].get('DB_HOST')
                self.db_name = o['id'].get('DB_NAME')
                self.db_username = o['id'].get('DB_USERNAME')
                self.db_pass = o['id'].get('DB_PASS')
                self.db_port = o['id'].get('DB_PORT')
