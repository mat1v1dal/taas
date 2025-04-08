from graphviz import Digraph

# Crear un nuevo diagrama con diseño más limpio y centrado
er = Digraph('ER_Improved', filename='taas_er_improved', format='png')
er.attr(rankdir='TB', fontsize='10', nodesep='0.5', ranksep='0.75')
er.attr('node', shape='plaintext', fontname='Helvetica')

# Función para crear nodos estilo tabla
def create_table(name, fields):
    label = f'''<<TABLE BORDER="1" CELLBORDER="1" CELLSPACING="0" CELLPADDING="4">
    <TR><TD BGCOLOR="#e0e0e0" COLSPAN="1"><B>{name}</B></TD></TR>'''
    for field in fields:
        label += f'<TR><TD ALIGN="LEFT">{field}</TD></TR>'
    label += '</TABLE>>'
    er.node(name, label=label)

# Definir tablas
create_table('users', [
    'id (PK)', 'name', 'email', 'password_hash', 'role', 'created_at'
])

create_table('emotions', [
    'id (PK)', 'user_id (FK)', 'mood', 'note', 'created_at'
])

create_table('chat_logs', [
    'id (PK)', 'user_id (FK)', 'sender', 'message', 'created_at'
])

create_table('sessions', [
    'id (PK)', 'user_id (FK)', 'therapist_id (FK)', 'scheduled_at', 'link', 'notes', 'created_at'
])

create_table('exercises', [
    'id (PK)', 'title', 'description', 'category', 'created_at'
])

create_table('user_exercises', [
    'id (PK)', 'user_id (FK)', 'exercise_id (FK)', 'status', 'completed_at'
])

# Definir relaciones
er.edge('users', 'emotions', label='1:N')
er.edge('users', 'chat_logs', label='1:N')
er.edge('users', 'sessions', label='1:N (as user)')
er.edge('users', 'sessions', label='1:N (as therapist)')
er.edge('users', 'user_exercises', label='1:N')
er.edge('exercises', 'user_exercises', label='1:N')

# Renderizar
er.render('./diagrama', format='png', cleanup=False)
'/mnt/data/taas_er_improved.png'