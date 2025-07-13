import pymysql
import click

@click.command()
@click.option('--host', default='127.0.0.1', help='MySQL host address')
@click.option('--user', default='crud_user', help='MySQL username')
@click.option('--password', default='crud_password', help='MySQL password')
@click.option('--db', default='crud_db', help='MySQL database name')
def mysqlconnect(host, user, password, db):
    """Test MySQL connection with configurable parameters"""
    # To connect MySQL database
    conn = pymysql.connect(
        host=host,
        user=user,
        password=password,
        db=db,
    )

    cur = conn.cursor()
    cur.execute("select @@version")
    output = cur.fetchall()
    print(output)

    # To close the connection
    conn.close()


# Driver Code
if __name__ == "__main__":
    mysqlconnect()
