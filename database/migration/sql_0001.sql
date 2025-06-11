DROP TABLE IF EXISTS public.sis_usuario_grupo;
DROP TABLE IF EXISTS public.sis_arvore_modulo;
DROP TABLE IF EXISTS public.sis_programa_arvore;
DROP TABLE IF EXISTS public.sis_usuario_programa;
DROP TABLE IF EXISTS public.sis_grupo_programa;
DROP TABLE IF EXISTS public.sis_produto;
DROP TABLE IF EXISTS public.sis_unidade_medida;
DROP TABLE IF EXISTS public.sis_categoria_produto;
DROP TABLE IF EXISTS public.sis_tipo_unidade_medida;
DROP TABLE IF EXISTS public.sis_modulo;
DROP TABLE IF EXISTS public.sis_arvore;
DROP TABLE IF EXISTS public.sis_programa;
DROP TABLE IF EXISTS public.sis_grupo;
DROP TABLE IF EXISTS public.sis_usuario;

CREATE TABLE public.sis_usuario (
	id_usuario              	BIGINT GENERATED ALWAYS AS IDENTITY (INCREMENT BY 1 MINVALUE 1 MAXVALUE 9223372036854775807 START 1 CACHE 1 NO CYCLE) NOT NULL
	, codigo_usuario			TEXT NOT NULL
	, nome_usuario				TEXT NOT NULL
	, email_usuario				TEXT NOT NULL
	, senha_usuario				BYTEA NOT NULL
	, cpf_cnpj					TEXT NOT NULL
	, uuid_usuario				UUID NOT NULL
	, data_inclusao				TIMESTAMP WITHOUT TIME ZONE NOT NULL
	, id_usuario_inclusao		BIGINT NOT NULL
	, data_alteracao			TIMESTAMP WITHOUT TIME ZONE
	, id_usuario_alteracao		BIGINT
	, data_exclusao				TIMESTAMP WITHOUT TIME ZONE
	, id_usuario_exclusao		BIGINT
	, registro_ativo			BOOLEAN DEFAULT TRUE
	, CONSTRAINT pk_sis_usuario_id_usuario PRIMARY KEY (id_usuario)
);

CREATE TABLE public.sis_grupo (
	id_grupo					BIGINT GENERATED ALWAYS AS IDENTITY (INCREMENT BY 1 MINVALUE 1 MAXVALUE 9223372036854775807 START 1 CACHE 1 NO CYCLE) NOT NULL
	, nome_grupo				TEXT NOT NULL
	, data_inclusao				TIMESTAMP WITHOUT TIME ZONE NOT NULL
	, id_usuario_inclusao		BIGINT NOT NULL
	, data_alteracao			TIMESTAMP WITHOUT TIME ZONE
	, id_usuario_alteracao		BIGINT
	, data_exclusao				TIMESTAMP WITHOUT TIME ZONE
	, id_usuario_exclusao		BIGINT
	, CONSTRAINT pk_sis_grupo_id_grupo PRIMARY KEY (id_grupo)
);

CREATE TABLE public.sis_usuario_grupo (
	id_usuario_grupo		BIGINT GENERATED ALWAYS AS IDENTITY (INCREMENT BY 1 MINVALUE 1 MAXVALUE 9223372036854775807 START 1 CACHE 1 NO CYCLE) NOT NULL
	, id_usuario			BIGINT NOT NULL
	, id_grupo				BIGINT NOT NULL
	, data_inclusao			TIMESTAMP WITHOUT TIME ZONE NOT NULL
	, id_usuario_inclusao	BIGINT NOT NULL
	, data_alteracao		TIMESTAMP WITHOUT TIME ZONE
	, id_usuario_alteracao	BIGINT
	, data_exclusao			TIMESTAMP WITHOUT TIME ZONE
	, id_usuario_exclusao	BIGINT
	, CONSTRAINT pk_sis_usuario_grupo_id_usuario_grupo PRIMARY KEY (id_usuario_grupo)
);

CREATE TABLE public.sis_produto (
	id_produto						BIGINT GENERATED ALWAYS AS IDENTITY (INCREMENT BY 1 MINVALUE 1 MAXVALUE 9223372036854775807 START 1 CACHE 1 NO CYCLE) NOT NULL
	, nome_produto					TEXT NOT NULL
	, codigo_produto				TEXT NOT NULL
	, uuid_produto					UUID NOT NULL
	, id_categoria_produto			BIGINT NOT NULL
	, id_unidade_medida_produto		BIGINT NOT NULL
	, observacao					TEXT
	, url_imagem_produto			TEXT
	, registro_ativo				BOOLEAN DEFAULT TRUE
	, data_inclusao					TIMESTAMP WITHOUT TIME ZONE NOT NULL
	, id_usuario_inclusao			BIGINT NOT NULL
	, data_alteracao				TIMESTAMP WITHOUT TIME ZONE
	, id_usuario_alteracao			BIGINT
	, data_exclusao					TIMESTAMP WITHOUT TIME ZONE
	, id_usuario_exclusao			BIGINT
	, CONSTRAINT pk_sis_produto_id_produto PRIMARY KEY (id_produto)
);

CREATE TABLE public.sis_categoria_produto (
	id_categoria_produto			BIGINT GENERATED ALWAYS AS IDENTITY (INCREMENT BY 1 MINVALUE 1 MAXVALUE 9223372036854775807 START 1 CACHE 1 NO CYCLE) NOT NULL
	, nome_categoria_produto		TEXT NOT NULL
	, codigo_categoria_produto		TEXT NOT NULL
	, uuid_categoria_produto		UUID NOT NULL
	, observacao					TEXT
	, registro_ativo				BOOLEAN DEFAULT TRUE
	, data_inclusao					TIMESTAMP WITHOUT TIME ZONE NOT NULL
	, id_usuario_inclusao			BIGINT NOT NULL
	, data_alteracao				TIMESTAMP WITHOUT TIME ZONE
	, id_usuario_alteracao			BIGINT
	, data_exclusao					TIMESTAMP WITHOUT TIME ZONE
	, id_usuario_exclusao			BIGINT
	, CONSTRAINT pk_sis_categoria_produto_id_categoria_produto PRIMARY KEY (id_categoria_produto)
);

CREATE TABLE public.sis_modulo (
	id_modulo					BIGINT GENERATED ALWAYS AS IDENTITY (INCREMENT BY 1 MINVALUE 1 MAXVALUE 9223372036854775807 START 1 CACHE 1 NO CYCLE) NOT NULL
	, nome_modulo				TEXT NOT NULL
	, url_icone_modulo			TEXT NOT NULL
	, descricao_modulo			TEXT NOT NULL
	, data_inclusao				TIMESTAMP WITHOUT TIME ZONE NOT NULL
	, id_usuario_inclusao		BIGINT NOT NULL
	, data_alteracao			TIMESTAMP WITHOUT TIME ZONE
	, id_usuario_alteracao		BIGINT
	, data_exclusao				TIMESTAMP WITHOUT TIME ZONE
	, id_usuario_exclusao		BIGINT
	, CONSTRAINT pk_sis_modulo_id_modulo PRIMARY KEY (id_modulo)
);

CREATE TABLE public.sis_arvore (
	id_arvore					BIGINT GENERATED ALWAYS AS IDENTITY (INCREMENT BY 1 MINVALUE 1 MAXVALUE 9223372036854775807 START 1 CACHE 1 NO CYCLE) NOT NULL
	, nome_arvore				TEXT NOT NULL
	, data_inclusao				TIMESTAMP WITHOUT TIME ZONE NOT NULL
	, id_usuario_inclusao		BIGINT NOT NULL
	, data_alteracao			TIMESTAMP WITHOUT TIME ZONE
	, id_usuario_alteracao		BIGINT
	, data_exclusao				TIMESTAMP WITHOUT TIME ZONE
	, id_usuario_exclusao		BIGINT
	, CONSTRAINT pk_sis_arvore_id_arvore PRIMARY KEY (id_arvore)
);

CREATE TABLE public.sis_arvore_modulo (
	id_arvore_modulo		BIGINT GENERATED ALWAYS AS IDENTITY (INCREMENT BY 1 MINVALUE 1 MAXVALUE 9223372036854775807 START 1 CACHE 1 NO CYCLE) NOT NULL
	, id_arvore				BIGINT NOT NULL
	, id_modulo				BIGINT NOT NULL
	, data_inclusao			TIMESTAMP WITHOUT TIME ZONE NOT NULL
	, id_usuario_inclusao	BIGINT NOT NULL
	, data_alteracao		TIMESTAMP WITHOUT TIME ZONE
	, id_usuario_alteracao	BIGINT
	, data_exclusao			TIMESTAMP WITHOUT TIME ZONE
	, id_usuario_exclusao	BIGINT
	, CONSTRAINT pk_sis_arvore_modulo_id_arvore_modulo PRIMARY KEY (id_arvore_modulo)
);

CREATE TABLE public.sis_programa (
	id_programa				BIGINT GENERATED ALWAYS AS IDENTITY (INCREMENT BY 1 MINVALUE 1 MAXVALUE 9223372036854775807 START 1 CACHE 1 NO CYCLE) NOT NULL
	, nome_programa			TEXT NOT NULL
	, data_inclusao			TIMESTAMP WITHOUT TIME ZONE NOT NULL
	, id_usuario_inclusao	BIGINT NOT NULL
	, data_alteracao		TIMESTAMP WITHOUT TIME ZONE
	, id_usuario_alteracao	BIGINT
	, data_exclusao			TIMESTAMP WITHOUT TIME ZONE
	, id_usuario_exclusao	BIGINT
	, CONSTRAINT pk_sis_programa_id_programa PRIMARY KEY (id_programa)
);

CREATE TABLE public.sis_programa_arvore (
	id_programa_arvore		BIGINT GENERATED ALWAYS AS IDENTITY (INCREMENT BY 1 MINVALUE 1 MAXVALUE 9223372036854775807 START 1 CACHE 1 NO CYCLE) NOT NULL
	, id_programa			BIGINT NOT NULL
	, id_arvore				BIGINT NOT NULL
	, data_inclusao			TIMESTAMP WITHOUT TIME ZONE NOT NULL
	, id_usuario_inclusao	BIGINT NOT NULL
	, data_alteracao		TIMESTAMP WITHOUT TIME ZONE
	, id_usuario_alteracao	BIGINT
	, data_exclusao			TIMESTAMP WITHOUT TIME ZONE
	, id_usuario_exclusao	BIGINT
	, CONSTRAINT pk_sis_programa_arvore_id_programa_arvore PRIMARY KEY (id_programa_arvore)
);

CREATE TABLE public.sis_usuario_programa (
	id_usuario_programa			BIGINT GENERATED ALWAYS AS IDENTITY (INCREMENT BY 1 MINVALUE 1 MAXVALUE 9223372036854775807 START 1 CACHE 1 NO CYCLE) NOT NULL
	, id_usuario				BIGINT NOT NULL
	, id_programa				BIGINT NOT NULL
	, data_inclusao				TIMESTAMP WITHOUT TIME ZONE NOT NULL
	, id_usuario_inclusao		BIGINT NOT NULL
	, data_alteracao			TIMESTAMP WITHOUT TIME ZONE
	, id_usuario_alteracao		BIGINT
	, data_exclusao				TIMESTAMP WITHOUT TIME ZONE
	, id_usuario_exclusao		BIGINT
	, CONSTRAINT pk_sis_usuario_programa_id_usuario_programa PRIMARY KEY (id_usuario_programa)
);

CREATE TABLE public.sis_grupo_programa (
	id_grupo_programa		BIGINT GENERATED ALWAYS AS IDENTITY (INCREMENT BY 1 MINVALUE 1 MAXVALUE 9223372036854775807 START 1 CACHE 1 NO CYCLE) NOT NULL
	, id_grupo				BIGINT NOT NULL
	, id_programa			BIGINT NOT NULL
	, data_inclusao			TIMESTAMP WITHOUT TIME ZONE NOT NULL
	, id_usuario_inclusao	BIGINT NOT NULL
	, data_alteracao		TIMESTAMP WITHOUT TIME ZONE
	, id_usuario_alteracao	BIGINT
	, data_exclusao			TIMESTAMP WITHOUT TIME ZONE
	, id_usuario_exclusao	BIGINT
	, CONSTRAINT pk_sis_grupo_programa_id_grupo_programa PRIMARY KEY (id_grupo_programa)
);

CREATE TABLE public.sis_unidade_medida (
	id_unidade_medida			BIGINT GENERATED ALWAYS AS IDENTITY (INCREMENT BY 1 MINVALUE 1 MAXVALUE 9223372036854775807 START 1 CACHE 1 NO CYCLE) NOT NULL
	, nome_unidade_medida		TEXT NOT NULL
	, codigo_unidade_medida		TEXT NOT NULL
	, id_tipo_unidade_medida	BIGINT NOT NULL
	, uuid_unidade_medida		UUID NOT NULL
	, observacao				TEXT
	, registro_ativo			BOOLEAN DEFAULT TRUE
	, data_inclusao				TIMESTAMP WITHOUT TIME ZONE NOT NULL
	, id_usuario_inclusao		BIGINT NOT NULL
	, data_alteracao			TIMESTAMP WITHOUT TIME ZONE
	, id_usuario_alteracao		BIGINT
	, data_exclusao				TIMESTAMP WITHOUT TIME ZONE
	, id_usuario_exclusao		BIGINT
	, CONSTRAINT pk_sis_unidade_medida_id_unidade_medida PRIMARY KEY (id_unidade_medida)
);

CREATE TABLE public.sis_tipo_unidade_medida (
	id_tipo_unidade_medida					BIGINT GENERATED ALWAYS AS IDENTITY (INCREMENT BY 1 MINVALUE 1 MAXVALUE 9223372036854775807 START 1 CACHE 1 NO CYCLE) NOT NULL
	, nome_tipo_unidade_medida				TEXT NOT NULL
	, codigo_tipo_unidade_medida			TEXT NOT NULL
	, uuid_tipo_unidade_medida				UUID NOT NULL
	, observacao							TEXT
	, registro_ativo						BOOLEAN DEFAULT TRUE
	, data_inclusao							TIMESTAMP WITHOUT TIME ZONE NOT NULL
	, id_usuario_inclusao					BIGINT NOT NULL
	, data_alteracao						TIMESTAMP WITHOUT TIME ZONE
	, id_usuario_alteracao					BIGINT
	, data_exclusao							TIMESTAMP WITHOUT TIME ZONE
	, id_usuario_exclusao					BIGINT
	, CONSTRAINT pk_sis_tipo_unidade_medida_id_tipo_unidade_medida PRIMARY KEY (id_tipo_unidade_medida)
);