--
-- PostgreSQL database dump
--

-- Dumped from database version 12.11
-- Dumped by pg_dump version 14.5 (Homebrew)

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

ALTER TABLE IF EXISTS ONLY public.air DROP CONSTRAINT IF EXISTS uniq_at;
DROP TABLE IF EXISTS public.air;
SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: air; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.air (
    at timestamp with time zone NOT NULL,
    score smallint NOT NULL,
    voc smallint NOT NULL,
    co2 smallint NOT NULL,
    temp real NOT NULL,
    pm25 smallint NOT NULL,
    humidity real NOT NULL,
    voc_index smallint NOT NULL,
    co2_index smallint NOT NULL,
    temp_index smallint NOT NULL,
    pm25_index smallint NOT NULL,
    humidity_index smallint NOT NULL
);


ALTER TABLE public.air OWNER TO postgres;

--
-- Name: air uniq_at; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.air
    ADD CONSTRAINT uniq_at UNIQUE (at);


--
-- Name: SCHEMA public; Type: ACL; Schema: -; Owner: postgres
--

REVOKE ALL ON SCHEMA public FROM rdsadmin;
REVOKE ALL ON SCHEMA public FROM PUBLIC;
GRANT ALL ON SCHEMA public TO postgres;
GRANT ALL ON SCHEMA public TO PUBLIC;


--
-- PostgreSQL database dump complete
--

