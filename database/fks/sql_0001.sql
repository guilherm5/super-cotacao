ALTER TABLE public.sis_tipo_unidade_medida
	ADD CONSTRAINT pk_sis_tipo_unidade_medida_id_tipo_unidade_medida PRIMARY KEY (id_tipo_unidade_medida);

ALTER TABLE public.sis_produto
	ADD CONSTRAINT fk_sis_produto_id_unidade_medida_produto
	FOREIGN KEY (id_unidade_medida_produto)
	REFERENCES public.sis_unidade_medida (id_unidade_medida);

ALTER TABLE public.sis_produto
	ADD CONSTRAINT fk_sis_produto_id_categoria_produto
	FOREIGN KEY (id_categoria_produt)
	REFERENCES public.sis_categoria_produto (id_categoria_produto);

ALTER TABLE public.sis_produto
	ADD CONSTRAINT fk_sis_produto_id_usuario_inclusao
	FOREIGN KEY (id_usuario_inclusao)
	REFERENCES public.sis_usuario (id_usuario);

ALTER TABLE public.sis_produto
	ADD CONSTRAINT fk_sis_produto_id_usuario_alteracao
	FOREIGN KEY (id_usuario_alteracao)
	REFERENCES public.sis_usuario (id_usuario);

ALTER TABLE public.sis_produto
	ADD CONSTRAINT fk_sis_produto_id_usuario_exclusao
	FOREIGN KEY (id_usuario_exclusao)
	REFERENCES public.sis_usuario (id_usuario);

ALTER TABLE public.sis_unidade_medida
	ADD CONSTRAINT fk_sis_unidade_medida_id_tipo_unidade_medida
	FOREIGN KEY (id_tipo_unidade_medida)
	REFERENCES public.sis_tipo_unidade_medida (id_tipo_unidade_medida);

ALTER TABLE public.sis_unidade_medida
	ADD CONSTRAINT fk_sis_unidade_medida_id_usuario_inclusao
	FOREIGN KEY (id_usuario_inclusao)
	REFERENCES public.sis_usuario (id_usuario);

ALTER TABLE public.sis_unidade_medida
	ADD CONSTRAINT fk_sis_unidade_medida_id_usuario_alteracao
	FOREIGN KEY (id_usuario_alteracao)
	REFERENCES public.sis_usuario (id_usuario);

ALTER TABLE public.sis_unidade_medida
	ADD CONSTRAINT fk_sis_unidade_medida_id_usuario_exclusao
	FOREIGN KEY (id_usuario_exclusao)
	REFERENCES public.sis_usuario (id_usuario);

ALTER TABLE public.sis_categoria_produto
	ADD CONSTRAINT fk_sis_categoria_produto_id_usuario_inclusao
	FOREIGN KEY (id_usuario_inclusao)
	REFERENCES public.sis_usuario (id_usuario);

ALTER TABLE public.sis_categoria_produto
	ADD CONSTRAINT fk_sis_categoria_produto_id_usuario_alteracao
	FOREIGN KEY (id_usuario_alteracao)
	REFERENCES public.sis_usuario (id_usuario);

ALTER TABLE public.sis_categoria_produto
	ADD CONSTRAINT fk_sis_categoria_produto_id_usuario_exclusao
	FOREIGN KEY (id_usuario_exclusao)
	REFERENCES public.sis_usuario (id_usuario);

ALTER TABLE public.sis_tipo_unidade_medida
	ADD CONSTRAINT fk_sis_tipo_unidade_medida_id_usuario_inclusao
	FOREIGN KEY (id_usuario_inclusao)
	REFERENCES public.sis_usuario (id_usuario);

ALTER TABLE public.sis_tipo_unidade_medida
	ADD CONSTRAINT fk_sis_tipo_unidade_medida_id_usuario_alteracao
	FOREIGN KEY (id_usuario_alteracao)
	REFERENCES public.sis_usuario (id_usuario);

ALTER TABLE public.sis_tipo_unidade_medida
	ADD CONSTRAINT fk_sis_tipo_unidade_medida_id_usuario_exclusao
	FOREIGN KEY (id_usuario_exclusao)
	REFERENCES public.sis_usuario (id_usuario);

ALTER TABLE public.sis_modulo
    ADD CONSTRAINT fk_sis_modulo_id_usuario_inclusao
    FOREIGN KEY (id_usuario_inclusao)
    REFERENCES public.sis_usuario (id_usuario);

ALTER TABLE public.sis_modulo
    ADD CONSTRAINT fk_sis_modulo_id_usuario_alteracao
    FOREIGN KEY (id_usuario_alteracao)
    REFERENCES public.sis_usuario (id_usuario);

ALTER TABLE public.sis_modulo
    ADD CONSTRAINT fk_sis_modulo_id_usuario_exclusao
    FOREIGN KEY (id_usuario_exclusao)
    REFERENCES public.sis_usuario (id_usuario);

ALTER TABLE public.sis_arvore
    ADD CONSTRAINT fk_sis_arvore_id_usuario_inclusao
    FOREIGN KEY (id_usuario_inclusao)
    REFERENCES public.sis_usuario (id_usuario);

ALTER TABLE public.sis_arvore
    ADD CONSTRAINT fk_sis_arvore_id_usuario_alteracao
    FOREIGN KEY (id_usuario_alteracao)
    REFERENCES public.sis_usuario (id_usuario);

ALTER TABLE public.sis_arvore
    ADD CONSTRAINT fk_sis_arvore_id_usuario_exclusao
    FOREIGN KEY (id_usuario_exclusao)
    REFERENCES public.sis_usuario (id_usuario);

ALTER TABLE public.sis_arvore_modulo
	ADD CONSTRAINT fk_sis_arvore_modulo_id_arvore
	FOREIGN KEY (id_arvore)
	REFERENCES public.sis_arvore (id_arvore);

ALTER TABLE public.sis_arvore_modulo
	ADD CONSTRAINT fk_sis_arvore_modulo_id_modulo
	FOREIGN KEY (id_modulo)
	REFERENCES public.sis_modulo (id_modulo);

ALTER TABLE public.sis_arvore_modulo
    ADD CONSTRAINT fk_sis_arvore_modulo_id_usuario_inclusao
    FOREIGN KEY (id_usuario_inclusao)
    REFERENCES public.sis_usuario (id_usuario);

ALTER TABLE public.sis_arvore_modulo
    ADD CONSTRAINT fk_sis_arvore_modulo_id_usuario_alteracao
    FOREIGN KEY (id_usuario_alteracao)
    REFERENCES public.sis_usuario (id_usuario);

ALTER TABLE public.sis_arvore_modulo
    ADD CONSTRAINT fk_sis_arvore_modulo_id_usuario_exclusao
    FOREIGN KEY (id_usuario_exclusao)
    REFERENCES public.sis_usuario (id_usuario);

ALTER TABLE public.sis_usuario_grupo
	ADD CONSTRAINT fk_sis_usuario_grupo_id_usuario
	FOREIGN KEY (id_usuario)
	REFERENCES public.sis_usuario (id_usuario);

ALTER TABLE public.sis_usuario_grupo
	ADD CONSTRAINT fk_sis_usuario_grupo_id_grupo
	FOREIGN KEY (id_grupo)
	REFERENCES public.sis_grupo (id_grupo);

ALTER TABLE public.sis_usuario_programa
	ADD CONSTRAINT fk_sis_usuario_programa_id_usuario
	FOREIGN KEY (id_usuario)
	REFERENCES public.sis_usuario (id_usuario);

ALTER TABLE public.sis_usuario_programa
	ADD CONSTRAINT fk_sis_usuario_programa_id_programa
	FOREIGN KEY (id_programa)
	REFERENCES public.sis_programa (id_programa);

ALTER TABLE public.sis_usuario_programa
    ADD CONSTRAINT fk_sis_usuario_programa_id_usuario_inclusao
    FOREIGN KEY (id_usuario_inclusao)
    REFERENCES public.sis_usuario (id_usuario);

ALTER TABLE public.sis_usuario_programa
    ADD CONSTRAINT fk_sis_usuario_programa_id_usuario_alteracao
    FOREIGN KEY (id_usuario_alteracao)
    REFERENCES public.sis_usuario (id_usuario);

ALTER TABLE public.sis_usuario_programa
    ADD CONSTRAINT fk_sis_usuario_programa_id_usuario_exclusao
    FOREIGN KEY (id_usuario_exclusao)
    REFERENCES public.sis_usuario (id_usuario);

ALTER TABLE public.sis_programa
    ADD CONSTRAINT fk_sis_programa_id_usuario_inclusao
    FOREIGN KEY (id_usuario_inclusao)
    REFERENCES public.sis_usuario (id_usuario);

ALTER TABLE public.sis_programa
    ADD CONSTRAINT fk_sis_programa_id_usuario_alteracao
    FOREIGN KEY (id_usuario_alteracao)
    REFERENCES public.sis_usuario (id_usuario);

ALTER TABLE public.sis_programa
    ADD CONSTRAINT fk_sis_programa_id_usuario_exclusao
    FOREIGN KEY (id_usuario_exclusao)
    REFERENCES public.sis_usuario (id_usuario);

ALTER TABLE public.sis_grupo_programa
	ADD CONSTRAINT fk_sis_grupo_programa_id_grupo
	FOREIGN KEY (id_grupo)
	REFERENCES public.sis_grupo (id_grupo);

ALTER TABLE public.sis_grupo_programa
	ADD CONSTRAINT fk_sis_grupo_programa_id_programa
	FOREIGN KEY (id_programa)
	REFERENCES public.sis_programa (id_programa);

ALTER TABLE public.sis_grupo_programa
    ADD CONSTRAINT fk_sis_grupo_programa_id_usuario_inclusao
    FOREIGN KEY (id_usuario_inclusao)
    REFERENCES public.sis_usuario (id_usuario);

ALTER TABLE public.sis_grupo_programa
    ADD CONSTRAINT fk_sis_grupo_programa_id_usuario_alteracao
    FOREIGN KEY (id_usuario_alteracao)
    REFERENCES public.sis_usuario (id_usuario);

ALTER TABLE public.sis_grupo_programa
    ADD CONSTRAINT fk_sis_grupo_programa_id_usuario_exclusao
    FOREIGN KEY (id_usuario_exclusao)
    REFERENCES public.sis_usuario (id_usuario);

ALTER TABLE public.sis_programa_arvore
	ADD CONSTRAINT fk_sis_programa_arvore_id_programa
	FOREIGN KEY (id_programa)
	REFERENCES public.sis_programa (id_programa);

ALTER TABLE public.sis_programa_arvore
	ADD CONSTRAINT fk_sis_programa_arvore_id_arvore
	FOREIGN KEY (id_arvore)
	REFERENCES public.sis_arvore (id_arvore);

ALTER TABLE public.sis_programa_arvore
    ADD CONSTRAINT fk_sis_programa_arvore_id_usuario_inclusao
    FOREIGN KEY (id_usuario_inclusao)
    REFERENCES public.sis_usuario (id_usuario);

ALTER TABLE public.sis_programa_arvore
    ADD CONSTRAINT fk_sis_programa_arvore_id_usuario_alteracao
    FOREIGN KEY (id_usuario_alteracao)
    REFERENCES public.sis_usuario (id_usuario);

ALTER TABLE public.sis_programa_arvore
    ADD CONSTRAINT fk_sis_programa_arvore_id_usuario_exclusao
    FOREIGN KEY (id_usuario_exclusao)
    REFERENCES public.sis_usuario (id_usuario);