--
-- PostgreSQL database dump
--

-- Dumped from database version 14.4 (Debian 14.4-1.pgdg110+1)
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

DROP TABLE IF EXISTS public.latest;
SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: latest; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.latest (
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


ALTER TABLE public.latest OWNER TO postgres;

--
-- PostgreSQL database dump complete
--

