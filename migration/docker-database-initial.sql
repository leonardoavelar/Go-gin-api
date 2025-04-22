create table usuarios(
    id serial primary key,
    nome varchar,
    profissao varchar
);

INSERT INTO usuarios(nome, profissao) VALUES
('Leonardo Avelar', 'Engenherio de Software'),
('Fernanda Kebleris', 'Psicologa');