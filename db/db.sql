--
-- PostgreSQL database dump
--

-- Dumped from database version 12.9 (Ubuntu 12.9-0ubuntu0.20.04.1)
-- Dumped by pg_dump version 12.9 (Ubuntu 12.9-0ubuntu0.20.04.1)

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
-- Name: movies; Type: TABLE; Schema: public; Owner: rizq
--

CREATE TABLE public.movies (
    id integer NOT NULL,
    title character varying,
    description character varying,
    release_year smallint
);


ALTER TABLE public.movies OWNER TO rizq;

--
-- Name: movies_id_seq; Type: SEQUENCE; Schema: public; Owner: rizq
--

CREATE SEQUENCE public.movies_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.movies_id_seq OWNER TO rizq;

--
-- Name: movies_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: rizq
--

ALTER SEQUENCE public.movies_id_seq OWNED BY public.movies.id;


--
-- Name: movies id; Type: DEFAULT; Schema: public; Owner: rizq
--

ALTER TABLE ONLY public.movies ALTER COLUMN id SET DEFAULT nextval('public.movies_id_seq'::regclass);


--
-- Data for Name: movies; Type: TABLE DATA; Schema: public; Owner: rizq
--

COPY public.movies (id, title, description, release_year) FROM stdin;
\.


--
-- Name: movies_id_seq; Type: SEQUENCE SET; Schema: public; Owner: rizq
--

SELECT pg_catalog.setval('public.movies_id_seq', 6, true);


--
-- Name: movies movies_pk; Type: CONSTRAINT; Schema: public; Owner: rizq
--

ALTER TABLE ONLY public.movies
    ADD CONSTRAINT movies_pk PRIMARY KEY (id);


--
-- PostgreSQL database dump complete
--

