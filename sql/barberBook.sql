CREATE TABLE IF NOT EXISTS cliente (
id serial PRIMARY KEY,
nome varchar(250) NOT NULL,
email varchar(250) NOT NULL UNIQUE,
contato varchar(20) UNIQUE NOT NULL,
senha varchar(100) NOT NULL,
data_criacao timestamp DEFAULT now() NOT NULL,
data_atualizacao timestamp
);

CREATE TABLE IF NOT EXISTS servico (
id serial PRIMARY KEY,
nome varchar(150) NOT NULL,
preco numeric(10, 2),
duracao interval,
data_criacao timestamp DEFAULT now() NOT NULL,
data_atualizacao timestamp
);

CREATE TABLE IF NOT EXISTS barbeiro (
id serial PRIMARY KEY,
nome varchar(150) NOT NULL,
contato varchar(20) UNIQUE,
data_criacao timestamp DEFAULT now() NOT NULL,
data_atualizacao timestamp
);

CREATE TABLE IF NOT EXISTS barbearia (
id serial PRIMARY KEY,
nome varchar(255) NOT NULL,
cidade varchar(150) NOT NULL,
rua varchar(255) NOT NULL,
numero_residencia integer NOT NULL,
ponto_referencia varchar(255),
contato varchar(20) UNIQUE,
data_criacao timestamp DEFAULT now() NOT NULL, 
data_atualizacao timestamp
);

CREATE TABLE IF NOT EXISTS horario_trabalho_barbeiro (
id serial PRIMARY KEY,
barbeiro_id integer NOT NULL,
dia_semana varchar(20) NOT NULL, 
horario_inicio time NOT NULL,
horario_almoco_inicio time NOT NULL,
horario_almoco_fim time NOT NULL,
horario_fim time NOT NULL,
data_criacao timestamp DEFAULT now() NOT NULL, 
data_atualizacao timestamp,
CONSTRAINT id_barbeiro_fk FOREIGN KEY (barbeiro_id) REFERENCES barbeiro(id)
);


CREATE TYPE tipo_status AS ENUM ('ativo', 'cancelado', 'pendente');

CREATE TABLE IF NOT EXISTS reserva (
    id serial PRIMARY KEY,
    barbeiro_id integer NOT NULL,
    cliente_id integer NOT NULL,
    barbearia_id integer NOT NULL,
    servico_id integer NOT NULL,
    data_reserva date NOT NULL,
    data_reserva_original date,
    horario_inicial_reserva time NOT NULL,
    status tipo_status DEFAULT 'ativo',
    horario_final time,
    data_criacao timestamp DEFAULT now() NOT NULL, 
    data_atualizacao timestamp,
    CONSTRAINT barbeiro_id_fk FOREIGN KEY (barbeiro_id) REFERENCES barbeiro(id),
    CONSTRAINT cliente_id_fk FOREIGN KEY (cliente_id) REFERENCES cliente(id),
    CONSTRAINT barbearia_id_fk FOREIGN KEY (barbearia_id) REFERENCES barbearia(id),
    CONSTRAINT servico_id_fk FOREIGN KEY (servico_id) REFERENCES servico(id)
);

CREATE TABLE IF NOT EXISTS horario_trabalho_excecao (
    id serial PRIMARY KEY,
    barbeiro_id integer NOT NULL,
    data_excecao date NOT NULL,
    motivo text,
    data_criacao timestamp DEFAULT now() NOT NULL,
    data_atualizacao timestamp,
    CONSTRAINT id_barbeiro_fk FOREIGN KEY (barbeiro_id) REFERENCES barbeiro(id)
);

CREATE OR REPLACE FUNCTION day_of_week_to_text(dow integer) RETURNS VARCHAR AS $$
BEGIN
    CASE dow
        WHEN 0 THEN RETURN 'domingo';
        WHEN 1 THEN RETURN 'segunda-feira';
        WHEN 2 THEN RETURN 'terça-feira';
        WHEN 3 THEN RETURN 'quarta-feira';
        WHEN 4 THEN RETURN 'quinta-feira';
        WHEN 5 THEN RETURN 'sexta-feira';
        WHEN 6 THEN RETURN 'sábado';
        ELSE RETURN NULL;
    END CASE;
END;
$$ LANGUAGE plpgsql;

-- Função trigger para validar o horário da reserva
CREATE OR REPLACE FUNCTION valida_horario_reserva() RETURNS TRIGGER AS $$
DECLARE
    dia_semana_reserva VARCHAR;
    duracao_servico interval;
BEGIN
    SELECT duracao INTO duracao_servico FROM servico WHERE id = NEW.servico_id;

    dia_semana_reserva := unaccent(day_of_week_to_text(EXTRACT(DOW FROM NEW.data_reserva::date)::integer));
    dia_semana_reserva := LOWER(REPLACE(TRIM(dia_semana_reserva), '-', ''));

    NEW.horario_final := NEW.horario_inicial_reserva + duracao_servico;
    -- Verifica se o dia da semana da reserva corresponde ao horário de trabalho do barbeiro
    IF EXISTS (
        SELECT 1
        FROM horario_trabalho_barbeiro
        WHERE barbeiro_id = NEW.barbeiro_id
        AND unaccent(LOWER(REPLACE(dia_semana, '-', ''))) = dia_semana_reserva -- Compara o dia da semana da reserva com o dia da semana especificado no horário do barbeiro
        AND NEW.horario_inicial_reserva >= horario_inicio
        AND (NEW.horario_inicial_reserva + duracao_servico) <= horario_fim
        AND (
            (NEW.horario_inicial_reserva >= horario_inicio AND (NEW.horario_inicial_reserva + duracao_servico) <= horario_almoco_inicio)
            OR
            (NEW.horario_inicial_reserva >= horario_almoco_fim AND (NEW.horario_inicial_reserva + duracao_servico) <= horario_fim)
        )
    ) THEN
        RETURN NEW; 
    ELSE
        RAISE EXCEPTION 'Reserva não permitida: Reserva fora do horário de trabalho do barbeiro para o dia da semana especificado';
    END IF;

EXCEPTION
    WHEN others THEN
        RAISE EXCEPTION 'Erro ao validar horário da reserva: %', SQLERRM;
END;
$$ LANGUAGE plpgsql;

DROP TRIGGER IF EXISTS trig_valida_horario_reserva ON reserva;

CREATE TRIGGER trig_valida_horario_reserva
BEFORE INSERT ON reserva
FOR EACH ROW EXECUTE FUNCTION valida_horario_reserva();

CREATE EXTENSION IF NOT EXISTS unaccent;
