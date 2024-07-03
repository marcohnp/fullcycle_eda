INSERT INTO clients (id, name, email, created_at)
VALUES ('85ef0de6-24f9-4f1a-a961-6514bfcdf80a', 'John Doe', 'john@j.com', '2024-02-13');
INSERT INTO clients (id, name, email, created_at)
VALUES ('dfced227-08a4-4594-9636-e6b3751ba6e1', 'Joanna Doe', 'joanna@j.com', '2024-02-13');

INSERT INTO accounts (id, client_id, balance, created_at)
VALUES ('206ef9f0-e240-4281-a4c8-6813e4f88861', 'dfced227-08a4-4594-9636-e6b3751ba6e1', 1000, '2024-07-01');
INSERT INTO accounts (id, client_id, balance, created_at)
VALUES ('3a2413c4-230a-4254-8744-f88764ad4b9a', '85ef0de6-24f9-4f1a-a961-6514bfcdf80a', 2000, '2024-07-01');
