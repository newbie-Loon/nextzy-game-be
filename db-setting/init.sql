CREATE TABLE public.point_histories (
	id bigserial NOT NULL,
	created_at timestamptz NULL,
	user_id varchar(255) NOT NULL,
	point int4 NOT NULL,
	earned_date timestamptz DEFAULT now() NOT NULL,
	deleted_at timestamptz NULL,
	updated_at timestamptz NULL,
	CONSTRAINT point_histories_pkey PRIMARY KEY (id)
);
CREATE INDEX idx_point_histories_earned_date ON public.point_histories USING btree (earned_date);
CREATE INDEX idx_point_histories_user_id ON public.point_histories USING btree (user_id);

CREATE TABLE public.reward_histories (
	id bigserial NOT NULL,
	created_at timestamptz NULL,
	user_id varchar(255) NOT NULL,
	reward varchar(255) NOT NULL,
	earned_date timestamptz DEFAULT now() NOT NULL,
	deleted_at timestamptz NULL,
	updated_at timestamptz NULL,
	CONSTRAINT reward_histories_pkey PRIMARY KEY (id)
);
CREATE INDEX idx_reward_histories_earned_date ON public.reward_histories USING btree (earned_date);
CREATE INDEX idx_reward_histories_user_id ON public.reward_histories USING btree (user_id);

CREATE TABLE public.users (
	id bigserial NOT NULL,
	created_at timestamptz NULL,
	updated_at timestamptz NULL,
	deleted_at timestamptz NULL,
	guest bool DEFAULT false NOT NULL,
	user_id varchar(255) DEFAULT gen_random_uuid() NOT NULL,
	CONSTRAINT users_pkey PRIMARY KEY (id)
);
CREATE INDEX idx_users_user_id ON public.users USING btree (user_id);