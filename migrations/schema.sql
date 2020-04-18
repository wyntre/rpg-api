--
-- PostgreSQL database dump
--

-- Dumped from database version 12.2
-- Dumped by pg_dump version 12.2

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: campaigns; Type: TABLE; Schema: public; Owner: buffalo
--

CREATE TABLE public.campaigns (
    id uuid NOT NULL,
    name character varying(255) NOT NULL,
    description character varying(255) NOT NULL,
    user_id uuid NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.campaigns OWNER TO buffalo;

--
-- Name: characters; Type: TABLE; Schema: public; Owner: buffalo
--

CREATE TABLE public.characters (
    id uuid NOT NULL,
    user_id uuid NOT NULL,
    name character varying(255) NOT NULL,
    description text NOT NULL,
    campaign_id uuid,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.characters OWNER TO buffalo;

--
-- Name: levels; Type: TABLE; Schema: public; Owner: buffalo
--

CREATE TABLE public.levels (
    id uuid NOT NULL,
    name character varying(255) NOT NULL,
    description character varying(255) NOT NULL,
    user_id uuid NOT NULL,
    map_id uuid NOT NULL,
    sort_order integer NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.levels OWNER TO buffalo;

--
-- Name: maps; Type: TABLE; Schema: public; Owner: buffalo
--

CREATE TABLE public.maps (
    id uuid NOT NULL,
    name character varying(255) NOT NULL,
    description character varying(255) NOT NULL,
    user_id uuid NOT NULL,
    quest_id uuid NOT NULL,
    sort_order integer NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.maps OWNER TO buffalo;

--
-- Name: quests; Type: TABLE; Schema: public; Owner: buffalo
--

CREATE TABLE public.quests (
    id uuid NOT NULL,
    name character varying(255) NOT NULL,
    description character varying(255) NOT NULL,
    user_id uuid NOT NULL,
    campaign_id uuid NOT NULL,
    sort_order integer NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.quests OWNER TO buffalo;

--
-- Name: revokedtokens; Type: TABLE; Schema: public; Owner: buffalo
--

CREATE TABLE public.revokedtokens (
    id uuid NOT NULL,
    token text NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.revokedtokens OWNER TO buffalo;

--
-- Name: schema_migration; Type: TABLE; Schema: public; Owner: buffalo
--

CREATE TABLE public.schema_migration (
    version character varying(14) NOT NULL
);


ALTER TABLE public.schema_migration OWNER TO buffalo;

--
-- Name: tile_categories; Type: TABLE; Schema: public; Owner: buffalo
--

CREATE TABLE public.tile_categories (
    id uuid NOT NULL,
    name character varying(255) NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.tile_categories OWNER TO buffalo;

--
-- Name: tile_types; Type: TABLE; Schema: public; Owner: buffalo
--

CREATE TABLE public.tile_types (
    id uuid NOT NULL,
    name character varying(255) NOT NULL,
    tile_category_id uuid NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.tile_types OWNER TO buffalo;

--
-- Name: tiles; Type: TABLE; Schema: public; Owner: buffalo
--

CREATE TABLE public.tiles (
    id uuid NOT NULL,
    user_id uuid NOT NULL,
    x integer NOT NULL,
    y integer NOT NULL,
    level_id uuid NOT NULL,
    tile_type_id uuid NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.tiles OWNER TO buffalo;

--
-- Name: users; Type: TABLE; Schema: public; Owner: buffalo
--

CREATE TABLE public.users (
    id uuid NOT NULL,
    email character varying(255) NOT NULL,
    password_hash character varying(255) NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.users OWNER TO buffalo;

--
-- Name: campaigns campaigns_pkey; Type: CONSTRAINT; Schema: public; Owner: buffalo
--

ALTER TABLE ONLY public.campaigns
    ADD CONSTRAINT campaigns_pkey PRIMARY KEY (id);


--
-- Name: characters characters_pkey; Type: CONSTRAINT; Schema: public; Owner: buffalo
--

ALTER TABLE ONLY public.characters
    ADD CONSTRAINT characters_pkey PRIMARY KEY (id);


--
-- Name: levels levels_pkey; Type: CONSTRAINT; Schema: public; Owner: buffalo
--

ALTER TABLE ONLY public.levels
    ADD CONSTRAINT levels_pkey PRIMARY KEY (id);


--
-- Name: maps maps_pkey; Type: CONSTRAINT; Schema: public; Owner: buffalo
--

ALTER TABLE ONLY public.maps
    ADD CONSTRAINT maps_pkey PRIMARY KEY (id);


--
-- Name: quests quests_pkey; Type: CONSTRAINT; Schema: public; Owner: buffalo
--

ALTER TABLE ONLY public.quests
    ADD CONSTRAINT quests_pkey PRIMARY KEY (id);


--
-- Name: revokedtokens revokedtokens_pkey; Type: CONSTRAINT; Schema: public; Owner: buffalo
--

ALTER TABLE ONLY public.revokedtokens
    ADD CONSTRAINT revokedtokens_pkey PRIMARY KEY (id);


--
-- Name: tile_categories tile_categories_pkey; Type: CONSTRAINT; Schema: public; Owner: buffalo
--

ALTER TABLE ONLY public.tile_categories
    ADD CONSTRAINT tile_categories_pkey PRIMARY KEY (id);


--
-- Name: tile_types tile_types_pkey; Type: CONSTRAINT; Schema: public; Owner: buffalo
--

ALTER TABLE ONLY public.tile_types
    ADD CONSTRAINT tile_types_pkey PRIMARY KEY (id);


--
-- Name: tiles tiles_pkey; Type: CONSTRAINT; Schema: public; Owner: buffalo
--

ALTER TABLE ONLY public.tiles
    ADD CONSTRAINT tiles_pkey PRIMARY KEY (id);


--
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: buffalo
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


--
-- Name: schema_migration_version_idx; Type: INDEX; Schema: public; Owner: buffalo
--

CREATE UNIQUE INDEX schema_migration_version_idx ON public.schema_migration USING btree (version);


--
-- Name: tiles_x_y_level_id_idx; Type: INDEX; Schema: public; Owner: buffalo
--

CREATE UNIQUE INDEX tiles_x_y_level_id_idx ON public.tiles USING btree (x, y, level_id);


--
-- Name: campaigns campaigns_user_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: buffalo
--

ALTER TABLE ONLY public.campaigns
    ADD CONSTRAINT campaigns_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.users(id);


--
-- Name: characters characters_campaign_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: buffalo
--

ALTER TABLE ONLY public.characters
    ADD CONSTRAINT characters_campaign_id_fkey FOREIGN KEY (campaign_id) REFERENCES public.campaigns(id);


--
-- Name: characters characters_user_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: buffalo
--

ALTER TABLE ONLY public.characters
    ADD CONSTRAINT characters_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.users(id);


--
-- Name: levels levels_map_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: buffalo
--

ALTER TABLE ONLY public.levels
    ADD CONSTRAINT levels_map_id_fkey FOREIGN KEY (map_id) REFERENCES public.maps(id);


--
-- Name: levels levels_user_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: buffalo
--

ALTER TABLE ONLY public.levels
    ADD CONSTRAINT levels_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.users(id);


--
-- Name: maps maps_quest_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: buffalo
--

ALTER TABLE ONLY public.maps
    ADD CONSTRAINT maps_quest_id_fkey FOREIGN KEY (quest_id) REFERENCES public.quests(id);


--
-- Name: maps maps_user_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: buffalo
--

ALTER TABLE ONLY public.maps
    ADD CONSTRAINT maps_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.users(id);


--
-- Name: quests quests_campaign_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: buffalo
--

ALTER TABLE ONLY public.quests
    ADD CONSTRAINT quests_campaign_id_fkey FOREIGN KEY (campaign_id) REFERENCES public.campaigns(id);


--
-- Name: quests quests_user_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: buffalo
--

ALTER TABLE ONLY public.quests
    ADD CONSTRAINT quests_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.users(id);


--
-- Name: tile_types tile_types_tile_category_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: buffalo
--

ALTER TABLE ONLY public.tile_types
    ADD CONSTRAINT tile_types_tile_category_id_fkey FOREIGN KEY (tile_category_id) REFERENCES public.tile_categories(id);


--
-- Name: tiles tiles_level_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: buffalo
--

ALTER TABLE ONLY public.tiles
    ADD CONSTRAINT tiles_level_id_fkey FOREIGN KEY (level_id) REFERENCES public.levels(id);


--
-- Name: tiles tiles_tile_type_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: buffalo
--

ALTER TABLE ONLY public.tiles
    ADD CONSTRAINT tiles_tile_type_id_fkey FOREIGN KEY (tile_type_id) REFERENCES public.tile_types(id);


--
-- Name: tiles tiles_user_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: buffalo
--

ALTER TABLE ONLY public.tiles
    ADD CONSTRAINT tiles_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.users(id);


--
-- PostgreSQL database dump complete
--

