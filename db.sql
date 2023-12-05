--
-- PostgreSQL database dump
--

-- Dumped from database version 14.9 (Homebrew)
-- Dumped by pg_dump version 14.9 (Homebrew)

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

DROP DATABASE IF EXISTS productsdb;
--
-- Name: productsdb; Type: DATABASE; Schema: -; Owner: postgres
--

CREATE DATABASE productsdb WITH TEMPLATE = template0 ENCODING = 'UTF8' LOCALE = 'C';


ALTER DATABASE productsdb OWNER TO postgres;

\connect productsdb

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

--
-- Name: public; Type: SCHEMA; Schema: -; Owner: pathaoltd
--

CREATE SCHEMA public;


ALTER SCHEMA public OWNER TO pathaoltd;

--
-- Name: SCHEMA public; Type: COMMENT; Schema: -; Owner: pathaoltd
--

COMMENT ON SCHEMA public IS 'standard public schema';


SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: brands; Type: TABLE; Schema: public; Owner: admin
--

CREATE TABLE public.brands (
                               id bigint NOT NULL,
                               created_at timestamp with time zone,
                               updated_at timestamp with time zone,
                               deleted_at timestamp with time zone,
                               name text,
                               status_id text
);


ALTER TABLE public.brands OWNER TO admin;

--
-- Name: brands_id_seq; Type: SEQUENCE; Schema: public; Owner: admin
--

CREATE SEQUENCE public.brands_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.brands_id_seq OWNER TO admin;

--
-- Name: brands_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: admin
--

ALTER SEQUENCE public.brands_id_seq OWNED BY public.brands.id;


--
-- Name: categories; Type: TABLE; Schema: public; Owner: admin
--

CREATE TABLE public.categories (
                                   id bigint NOT NULL,
                                   created_at timestamp with time zone,
                                   updated_at timestamp with time zone,
                                   deleted_at timestamp with time zone,
                                   name text,
                                   parent_id bigint,
                                   sequence bigint,
                                   status_id text
);


ALTER TABLE public.categories OWNER TO admin;

--
-- Name: categories_id_seq; Type: SEQUENCE; Schema: public; Owner: admin
--

CREATE SEQUENCE public.categories_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.categories_id_seq OWNER TO admin;

--
-- Name: categories_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: admin
--

ALTER SEQUENCE public.categories_id_seq OWNED BY public.categories.id;


--
-- Name: products; Type: TABLE; Schema: public; Owner: admin
--

CREATE TABLE public.products (
                                 id bigint NOT NULL,
                                 created_at timestamp with time zone,
                                 updated_at timestamp with time zone,
                                 deleted_at timestamp with time zone,
                                 name text,
                                 description character varying(100),
                                 specifications text,
                                 unit_price bigint,
                                 discount_price bigint,
                                 tags text,
                                 status_id text,
                                 brand_id bigint,
                                 category_id bigint,
                                 supplier_id bigint
);


ALTER TABLE public.products OWNER TO admin;

--
-- Name: products_id_seq; Type: SEQUENCE; Schema: public; Owner: admin
--

CREATE SEQUENCE public.products_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.products_id_seq OWNER TO admin;

--
-- Name: products_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: admin
--

ALTER SEQUENCE public.products_id_seq OWNED BY public.products.id;


--
-- Name: products_stocks; Type: TABLE; Schema: public; Owner: admin
--

CREATE TABLE public.products_stocks (
                                        id bigint NOT NULL,
                                        created_at timestamp with time zone,
                                        updated_at timestamp with time zone,
                                        deleted_at timestamp with time zone,
                                        products_id bigint,
                                        stock_quantity bigint
);


ALTER TABLE public.products_stocks OWNER TO admin;

--
-- Name: products_stocks_id_seq; Type: SEQUENCE; Schema: public; Owner: admin
--

CREATE SEQUENCE public.products_stocks_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.products_stocks_id_seq OWNER TO admin;

--
-- Name: products_stocks_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: admin
--

ALTER SEQUENCE public.products_stocks_id_seq OWNED BY public.products_stocks.id;


--
-- Name: suppliers; Type: TABLE; Schema: public; Owner: admin
--

CREATE TABLE public.suppliers (
                                  id bigint NOT NULL,
                                  created_at timestamp with time zone,
                                  updated_at timestamp with time zone,
                                  deleted_at timestamp with time zone,
                                  name text,
                                  email text,
                                  phone text,
                                  status_id text,
                                  is_verified_supplier boolean
);


ALTER TABLE public.suppliers OWNER TO admin;

--
-- Name: suppliers_id_seq; Type: SEQUENCE; Schema: public; Owner: admin
--

CREATE SEQUENCE public.suppliers_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.suppliers_id_seq OWNER TO admin;

--
-- Name: suppliers_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: admin
--

ALTER SEQUENCE public.suppliers_id_seq OWNED BY public.suppliers.id;


--
-- Name: brands id; Type: DEFAULT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.brands ALTER COLUMN id SET DEFAULT nextval('public.brands_id_seq'::regclass);


--
-- Name: categories id; Type: DEFAULT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.categories ALTER COLUMN id SET DEFAULT nextval('public.categories_id_seq'::regclass);


--
-- Name: products id; Type: DEFAULT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.products ALTER COLUMN id SET DEFAULT nextval('public.products_id_seq'::regclass);


--
-- Name: products_stocks id; Type: DEFAULT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.products_stocks ALTER COLUMN id SET DEFAULT nextval('public.products_stocks_id_seq'::regclass);


--
-- Name: suppliers id; Type: DEFAULT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.suppliers ALTER COLUMN id SET DEFAULT nextval('public.suppliers_id_seq'::regclass);


--
-- Data for Name: brands; Type: TABLE DATA; Schema: public; Owner: admin
--

INSERT INTO public.brands (id, created_at, updated_at, deleted_at, name, status_id) VALUES (23, '2023-12-05 12:59:33.647644+06', '2023-12-05 12:59:33.647644+06', NULL, 'Fog', '1');
INSERT INTO public.brands (id, created_at, updated_at, deleted_at, name, status_id) VALUES (12, '2023-12-05 12:59:29.31001+06', '2023-12-05 12:59:29.31001+06', NULL, 'Walton', '1');
INSERT INTO public.brands (id, created_at, updated_at, deleted_at, name, status_id) VALUES (15, '2023-12-05 12:59:30.422262+06', '2023-12-05 12:59:30.422262+06', NULL, 'La Reve', '1');
INSERT INTO public.brands (id, created_at, updated_at, deleted_at, name, status_id) VALUES (9, '2023-12-05 12:59:28.066057+06', '2023-12-05 12:59:28.066057+06', NULL, 'Apple', '1');
INSERT INTO public.brands (id, created_at, updated_at, deleted_at, name, status_id) VALUES (8, '2023-12-05 12:59:27.671463+06', '2023-12-05 12:59:27.671463+06', NULL, 'Garnier', '1');
INSERT INTO public.brands (id, created_at, updated_at, deleted_at, name, status_id) VALUES (14, '2023-12-05 12:59:30.059227+06', '2023-12-05 12:59:30.059227+06', NULL, 'Aarong', '1');
INSERT INTO public.brands (id, created_at, updated_at, deleted_at, name, status_id) VALUES (1, '2023-12-05 12:59:22.079821+06', '2023-12-05 12:59:22.079821+06', NULL, 'RFL', '1');
INSERT INTO public.brands (id, created_at, updated_at, deleted_at, name, status_id) VALUES (11, '2023-12-05 12:59:28.89989+06', '2023-12-05 12:59:28.89989+06', NULL, 'Hitachi', '1');
INSERT INTO public.brands (id, created_at, updated_at, deleted_at, name, status_id) VALUES (2, '2023-12-05 12:59:24.030961+06', '2023-12-05 12:59:24.030961+06', NULL, 'President', '1');
INSERT INTO public.brands (id, created_at, updated_at, deleted_at, name, status_id) VALUES (19, '2023-12-05 12:59:31.965368+06', '2023-12-05 12:59:31.965368+06', NULL, 'Abaya', '1');
INSERT INTO public.brands (id, created_at, updated_at, deleted_at, name, status_id) VALUES (18, '2023-12-05 12:59:31.577555+06', '2023-12-05 12:59:31.577555+06', NULL, 'Canon', '1');
INSERT INTO public.brands (id, created_at, updated_at, deleted_at, name, status_id) VALUES (7, '2023-12-05 12:59:27.21119+06', '2023-12-05 12:59:27.21119+06', NULL, 'Nestle', '1');
INSERT INTO public.brands (id, created_at, updated_at, deleted_at, name, status_id) VALUES (20, '2023-12-05 12:59:32.435453+06', '2023-12-05 12:59:32.435453+06', NULL, 'Curren', '1');
INSERT INTO public.brands (id, created_at, updated_at, deleted_at, name, status_id) VALUES (10, '2023-12-05 12:59:28.488623+06', '2023-12-05 12:59:28.488623+06', NULL, 'Gucci', '1');
INSERT INTO public.brands (id, created_at, updated_at, deleted_at, name, status_id) VALUES (6, '2023-12-05 12:59:26.790028+06', '2023-12-05 12:59:26.790028+06', NULL, 'Bata', '1');
INSERT INTO public.brands (id, created_at, updated_at, deleted_at, name, status_id) VALUES (5, '2023-12-05 12:59:26.349407+06', '2023-12-05 12:59:26.349407+06', NULL, 'Nokia', '1');
INSERT INTO public.brands (id, created_at, updated_at, deleted_at, name, status_id) VALUES (21, '2023-12-05 12:59:32.764558+06', '2023-12-05 12:59:32.764558+06', NULL, 'Godrej', '1');
INSERT INTO public.brands (id, created_at, updated_at, deleted_at, name, status_id) VALUES (3, '2023-12-05 12:59:25.459693+06', '2023-12-05 12:59:25.459693+06', NULL, 'Milan', '1');
INSERT INTO public.brands (id, created_at, updated_at, deleted_at, name, status_id) VALUES (22, '2023-12-05 12:59:33.222047+06', '2023-12-05 12:59:33.222047+06', NULL, 'Lakme', '1');
INSERT INTO public.brands (id, created_at, updated_at, deleted_at, name, status_id) VALUES (17, '2023-12-05 12:59:31.209702+06', '2023-12-05 12:59:31.209702+06', NULL, 'Samsung', '1');
INSERT INTO public.brands (id, created_at, updated_at, deleted_at, name, status_id) VALUES (4, '2023-12-05 12:59:25.873957+06', '2023-12-05 12:59:25.873957+06', NULL, 'Adidas', '1');
INSERT INTO public.brands (id, created_at, updated_at, deleted_at, name, status_id) VALUES (16, '2023-12-05 12:59:30.868041+06', '2023-12-05 12:59:30.868041+06', NULL, 'Dell', '1');
INSERT INTO public.brands (id, created_at, updated_at, deleted_at, name, status_id) VALUES (13, '2023-12-05 12:59:29.680113+06', '2023-12-05 12:59:29.680113+06', NULL, 'Redmi', '1');
INSERT INTO public.brands (id, created_at, updated_at, deleted_at, name, status_id) VALUES (24, '2023-12-05 13:01:27.045158+06', '2023-12-05 13:01:27.045158+06', NULL, 'Prefunmance', '1');


--
-- Data for Name: categories; Type: TABLE DATA; Schema: public; Owner: admin
--

INSERT INTO public.categories (id, created_at, updated_at, deleted_at, name, parent_id, sequence, status_id) VALUES (1, '2023-12-05 12:37:46.547674+06', '2023-12-05 12:37:46.547674+06', NULL, 'Electronics', NULL, 1, '1');
INSERT INTO public.categories (id, created_at, updated_at, deleted_at, name, parent_id, sequence, status_id) VALUES (2, '2023-12-05 12:38:55.369322+06', '2023-12-05 12:38:55.369322+06', NULL, 'Cloths', NULL, 2, '1');
INSERT INTO public.categories (id, created_at, updated_at, deleted_at, name, parent_id, sequence, status_id) VALUES (3, '2023-12-05 12:39:04.699255+06', '2023-12-05 12:39:04.699255+06', NULL, 'Fashions', NULL, 3, '1');
INSERT INTO public.categories (id, created_at, updated_at, deleted_at, name, parent_id, sequence, status_id) VALUES (4, '2023-12-05 12:39:24.840937+06', '2023-12-05 12:39:24.840937+06', NULL, 'men', 2, 1, '1');
INSERT INTO public.categories (id, created_at, updated_at, deleted_at, name, parent_id, sequence, status_id) VALUES (5, '2023-12-05 12:39:28.257451+06', '2023-12-05 12:39:28.257451+06', NULL, 'men', 3, 1, '1');
INSERT INTO public.categories (id, created_at, updated_at, deleted_at, name, parent_id, sequence, status_id) VALUES (6, '2023-12-05 12:40:44.42776+06', '2023-12-05 12:40:44.42776+06', NULL, 'women', 2, 2, '1');
INSERT INTO public.categories (id, created_at, updated_at, deleted_at, name, parent_id, sequence, status_id) VALUES (7, '2023-12-05 12:40:49.599608+06', '2023-12-05 12:40:49.599608+06', NULL, 'women', 3, 2, '1');
INSERT INTO public.categories (id, created_at, updated_at, deleted_at, name, parent_id, sequence, status_id) VALUES (8, '2023-12-05 12:41:23.046723+06', '2023-12-05 12:41:23.046723+06', NULL, 'Traditional', 6, 1, '1');
INSERT INTO public.categories (id, created_at, updated_at, deleted_at, name, parent_id, sequence, status_id) VALUES (9, '2023-12-05 12:42:18.679266+06', '2023-12-05 12:42:18.679266+06', NULL, 'Western', 6, 2, '1');
INSERT INTO public.categories (id, created_at, updated_at, deleted_at, name, parent_id, sequence, status_id) VALUES (10, '2023-12-05 12:42:26.082527+06', '2023-12-05 12:42:26.082527+06', NULL, 'Muslim', 6, 3, '1');
INSERT INTO public.categories (id, created_at, updated_at, deleted_at, name, parent_id, sequence, status_id) VALUES (11, '2023-12-05 12:43:10.838772+06', '2023-12-05 12:43:10.838772+06', NULL, 'Shoes', 7, 1, '1');
INSERT INTO public.categories (id, created_at, updated_at, deleted_at, name, parent_id, sequence, status_id) VALUES (12, '2023-12-05 12:43:17.693832+06', '2023-12-05 12:43:17.693832+06', NULL, 'Bags', 7, 2, '1');
INSERT INTO public.categories (id, created_at, updated_at, deleted_at, name, parent_id, sequence, status_id) VALUES (13, '2023-12-05 12:43:51.066264+06', '2023-12-05 12:43:51.066264+06', NULL, 'Watches', 7, 3, '1');
INSERT INTO public.categories (id, created_at, updated_at, deleted_at, name, parent_id, sequence, status_id) VALUES (14, '2023-12-05 12:44:38.771345+06', '2023-12-05 12:44:38.771345+06', NULL, 'Beauty', NULL, 4, '1');
INSERT INTO public.categories (id, created_at, updated_at, deleted_at, name, parent_id, sequence, status_id) VALUES (15, '2023-12-05 12:45:04.716335+06', '2023-12-05 12:45:04.716335+06', NULL, 'Skincare', 14, 1, '1');
INSERT INTO public.categories (id, created_at, updated_at, deleted_at, name, parent_id, sequence, status_id) VALUES (16, '2023-12-05 12:45:13.373986+06', '2023-12-05 12:45:13.373986+06', NULL, 'Haircare', 14, 2, '1');
INSERT INTO public.categories (id, created_at, updated_at, deleted_at, name, parent_id, sequence, status_id) VALUES (17, '2023-12-05 12:45:21.668661+06', '2023-12-05 12:45:21.668661+06', NULL, 'Makeup', 14, 3, '1');
INSERT INTO public.categories (id, created_at, updated_at, deleted_at, name, parent_id, sequence, status_id) VALUES (18, '2023-12-05 12:45:31.613055+06', '2023-12-05 12:45:31.613055+06', NULL, 'Fragrance', 14, 4, '1');
INSERT INTO public.categories (id, created_at, updated_at, deleted_at, name, parent_id, sequence, status_id) VALUES (19, '2023-12-05 12:46:18.305507+06', '2023-12-05 12:46:18.305507+06', NULL, 'Phone', 1, 1, '1');
INSERT INTO public.categories (id, created_at, updated_at, deleted_at, name, parent_id, sequence, status_id) VALUES (20, '2023-12-05 12:46:32.077787+06', '2023-12-05 12:46:32.077787+06', NULL, 'Televison', 1, 2, '1');
INSERT INTO public.categories (id, created_at, updated_at, deleted_at, name, parent_id, sequence, status_id) VALUES (21, '2023-12-05 12:46:46.454594+06', '2023-12-05 12:46:46.454594+06', NULL, 'Laptop', 1, 3, '1');
INSERT INTO public.categories (id, created_at, updated_at, deleted_at, name, parent_id, sequence, status_id) VALUES (22, '2023-12-05 12:46:56.117043+06', '2023-12-05 12:46:56.117043+06', NULL, 'Camera', 1, 4, '1');


--
-- Data for Name: products; Type: TABLE DATA; Schema: public; Owner: admin
--

INSERT INTO public.products (id, created_at, updated_at, deleted_at, name, description, specifications, unit_price, discount_price, tags, status_id, brand_id, category_id, supplier_id) VALUES (1, '2023-12-05 13:24:06.638668+06', '2023-12-05 13:24:06.638668+06', NULL, 'White Fatua', 'Lorem ipsum dolor sit amet, consectetur adipiscing elit', 'specifications details', 1700, 1500, 'white,fatua', '1', 14, 4, 3);
INSERT INTO public.products (id, created_at, updated_at, deleted_at, name, description, specifications, unit_price, discount_price, tags, status_id, brand_id, category_id, supplier_id) VALUES (4, '2023-12-05 13:57:37.565666+06', '2023-12-05 13:57:37.565666+06', NULL, 'Black Fatua', 'Lorem ipsum dolor sit amet, consectetur adipiscing elit', 'specifications details', 3700, 2500, 'white,fatua', '1', 14, 4, 3);
INSERT INTO public.products (id, created_at, updated_at, deleted_at, name, description, specifications, unit_price, discount_price, tags, status_id, brand_id, category_id, supplier_id) VALUES (3, '2023-12-05 13:40:48.387281+06', '2023-12-05 13:40:48.387281+06', NULL, 'Black Fatua', 'Lorem ipsum dolor sit amet, consectetur adipiscing elit', 'specifications details', 3700, 2500, 'white,fatua', '1', 14, 4, 3);
INSERT INTO public.products (id, created_at, updated_at, deleted_at, name, description, specifications, unit_price, discount_price, tags, status_id, brand_id, category_id, supplier_id) VALUES (2, '2023-12-05 13:34:17.400413+06', '2023-12-05 14:17:39.484637+06', NULL, 'Blue Fatua', 'Lorem ipsum dolor sit amet, consectetur adipiscing elit', 'specifications details', 1700, 1600, 'blue,fatua', '1', 14, 4, 3);


--
-- Data for Name: products_stocks; Type: TABLE DATA; Schema: public; Owner: admin
--



--
-- Data for Name: suppliers; Type: TABLE DATA; Schema: public; Owner: admin
--

INSERT INTO public.suppliers (id, created_at, updated_at, deleted_at, name, email, phone, status_id, is_verified_supplier) VALUES (1, '2023-12-05 12:05:48.524993+06', '2023-12-05 12:05:58.075135+06', NULL, 'Walton', 'walton@waltonbd.com', '01912121212', '1', true);
INSERT INTO public.suppliers (id, created_at, updated_at, deleted_at, name, email, phone, status_id, is_verified_supplier) VALUES (2, '2023-12-05 12:06:43.669956+06', '2023-12-05 12:06:43.669956+06', NULL, 'Global Brand', 'globalbrand@gmail.com', '01912121212', '1', true);
INSERT INTO public.suppliers (id, created_at, updated_at, deleted_at, name, email, phone, status_id, is_verified_supplier) VALUES (3, '2023-12-05 12:07:10.645856+06', '2023-12-05 12:07:10.645856+06', NULL, 'Aarong', 'aarong71@gmail.com', '01912121212', '1', true);
INSERT INTO public.suppliers (id, created_at, updated_at, deleted_at, name, email, phone, status_id, is_verified_supplier) VALUES (4, '2023-12-05 12:07:33.389709+06', '2023-12-05 12:07:33.389709+06', NULL, 'RFL', 'pran-rfl@gmail.com', '01912121212', '1', true);
INSERT INTO public.suppliers (id, created_at, updated_at, deleted_at, name, email, phone, status_id, is_verified_supplier) VALUES (5, '2023-12-05 12:08:18.352128+06', '2023-12-05 12:08:18.352128+06', NULL, 'Milan', 'milan22@gmail.com', '01912121212', '1', false);
INSERT INTO public.suppliers (id, created_at, updated_at, deleted_at, name, email, phone, status_id, is_verified_supplier) VALUES (6, '2023-12-05 12:08:38.685743+06', '2023-12-05 12:08:38.685743+06', NULL, 'President', 'presidentbag@gmail.com', '01912121212', '1', false);
INSERT INTO public.suppliers (id, created_at, updated_at, deleted_at, name, email, phone, status_id, is_verified_supplier) VALUES (7, '2023-12-05 12:09:47.841756+06', '2023-12-05 12:09:47.841756+06', NULL, 'Polar', 'polarbd@gmail.com', '01912121212', '1', false);
INSERT INTO public.suppliers (id, created_at, updated_at, deleted_at, name, email, phone, status_id, is_verified_supplier) VALUES (8, '2023-12-05 12:10:01.486597+06', '2023-12-05 12:10:01.486597+06', NULL, 'Nestle', 'nestle@gmail.com', '01912121212', '1', true);
INSERT INTO public.suppliers (id, created_at, updated_at, deleted_at, name, email, phone, status_id, is_verified_supplier) VALUES (9, '2023-12-05 12:10:44.919399+06', '2023-12-05 12:10:44.919399+06', NULL, 'Shoe Arena', 'shoearenabd@gmail.com', '01912121212', '1', false);
INSERT INTO public.suppliers (id, created_at, updated_at, deleted_at, name, email, phone, status_id, is_verified_supplier) VALUES (10, '2023-12-05 12:11:15.153549+06', '2023-12-05 12:11:15.153549+06', NULL, 'Shajgoj', 'shajbd@gmail.com', '01912121212', '1', true);
INSERT INTO public.suppliers (id, created_at, updated_at, deleted_at, name, email, phone, status_id, is_verified_supplier) VALUES (11, '2023-12-05 12:12:02.965204+06', '2023-12-05 12:12:02.965204+06', NULL, 'Gadgets and Gear', 'gandg@gmail.com', '01912121212', '1', true);
INSERT INTO public.suppliers (id, created_at, updated_at, deleted_at, name, email, phone, status_id, is_verified_supplier) VALUES (12, '2023-12-05 12:13:09.310206+06', '2023-12-05 12:13:09.310206+06', NULL, 'Beauty Bazar', 'beautybazar@gmail.com', '01912121212', '1', false);
INSERT INTO public.suppliers (id, created_at, updated_at, deleted_at, name, email, phone, status_id, is_verified_supplier) VALUES (13, '2023-12-05 13:05:04.863497+06', '2023-12-05 13:05:04.863497+06', NULL, 'Godrej Int', 'godrejco@gmail.com', '01912121212', '1', false);
INSERT INTO public.suppliers (id, created_at, updated_at, deleted_at, name, email, phone, status_id, is_verified_supplier) VALUES (14, '2023-12-05 13:05:59.937785+06', '2023-12-05 13:05:59.937785+06', NULL, 'Face Falls', 'faceeco@gmail.com', '01912121212', '1', false);
INSERT INTO public.suppliers (id, created_at, updated_at, deleted_at, name, email, phone, status_id, is_verified_supplier) VALUES (15, '2023-12-05 13:06:42.992058+06', '2023-12-05 13:06:42.992058+06', NULL, 'Gucci', 'guccibd@gmail.com', '01912121212', '1', false);
INSERT INTO public.suppliers (id, created_at, updated_at, deleted_at, name, email, phone, status_id, is_verified_supplier) VALUES (17, '2023-12-05 13:08:02.677358+06', '2023-12-05 13:08:02.677358+06', NULL, 'Abaya Origin', 'obborg@gmail.com', '01912121212', '1', false);
INSERT INTO public.suppliers (id, created_at, updated_at, deleted_at, name, email, phone, status_id, is_verified_supplier) VALUES (16, '2023-12-05 13:07:35.5693+06', '2023-12-05 13:09:25.555445+06', NULL, 'Xiaomi', 'xiaomebd@gmail.com', '01912121212', '1', true);
INSERT INTO public.suppliers (id, created_at, updated_at, deleted_at, name, email, phone, status_id, is_verified_supplier) VALUES (18, '2023-12-05 13:09:56.761252+06', '2023-12-05 13:09:56.761252+06', NULL, 'CurrenBd', 'currenwatch@gmail.com', '01912121212', '1', false);
INSERT INTO public.suppliers (id, created_at, updated_at, deleted_at, name, email, phone, status_id, is_verified_supplier) VALUES (19, '2023-12-05 13:11:26.922538+06', '2023-12-05 13:11:26.922538+06', NULL, 'Camera Home', 'camerahome@gmail.com', '01912121212', '1', false);
INSERT INTO public.suppliers (id, created_at, updated_at, deleted_at, name, email, phone, status_id, is_verified_supplier) VALUES (20, '2023-12-05 13:12:12.443545+06', '2023-12-05 13:12:12.443545+06', NULL, 'Addidas', 'addidasint@gmail.com', '01912121212', '1', true);


--
-- Name: brands_id_seq; Type: SEQUENCE SET; Schema: public; Owner: admin
--

SELECT pg_catalog.setval('public.brands_id_seq', 24, true);


--
-- Name: categories_id_seq; Type: SEQUENCE SET; Schema: public; Owner: admin
--

SELECT pg_catalog.setval('public.categories_id_seq', 22, true);


--
-- Name: products_id_seq; Type: SEQUENCE SET; Schema: public; Owner: admin
--

SELECT pg_catalog.setval('public.products_id_seq', 4, true);


--
-- Name: products_stocks_id_seq; Type: SEQUENCE SET; Schema: public; Owner: admin
--

SELECT pg_catalog.setval('public.products_stocks_id_seq', 1, false);


--
-- Name: suppliers_id_seq; Type: SEQUENCE SET; Schema: public; Owner: admin
--

SELECT pg_catalog.setval('public.suppliers_id_seq', 20, true);


--
-- Name: brands brands_pkey; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.brands
    ADD CONSTRAINT brands_pkey PRIMARY KEY (id);


--
-- Name: categories categories_pkey; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.categories
    ADD CONSTRAINT categories_pkey PRIMARY KEY (id);


--
-- Name: products products_pkey; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.products
    ADD CONSTRAINT products_pkey PRIMARY KEY (id);


--
-- Name: products_stocks products_stocks_pkey; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.products_stocks
    ADD CONSTRAINT products_stocks_pkey PRIMARY KEY (id);


--
-- Name: suppliers suppliers_pkey; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.suppliers
    ADD CONSTRAINT suppliers_pkey PRIMARY KEY (id);


--
-- Name: idx_brands_deleted_at; Type: INDEX; Schema: public; Owner: admin
--

CREATE INDEX idx_brands_deleted_at ON public.brands USING btree (deleted_at);


--
-- Name: idx_categories_deleted_at; Type: INDEX; Schema: public; Owner: admin
--

CREATE INDEX idx_categories_deleted_at ON public.categories USING btree (deleted_at);


--
-- Name: idx_products_deleted_at; Type: INDEX; Schema: public; Owner: admin
--

CREATE INDEX idx_products_deleted_at ON public.products USING btree (deleted_at);


--
-- Name: idx_products_stocks_deleted_at; Type: INDEX; Schema: public; Owner: admin
--

CREATE INDEX idx_products_stocks_deleted_at ON public.products_stocks USING btree (deleted_at);


--
-- Name: idx_suppliers_deleted_at; Type: INDEX; Schema: public; Owner: admin
--

CREATE INDEX idx_suppliers_deleted_at ON public.suppliers USING btree (deleted_at);


--
-- Name: categories fk_categories_sub_categories; Type: FK CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.categories
    ADD CONSTRAINT fk_categories_sub_categories FOREIGN KEY (parent_id) REFERENCES public.categories(id);


--
-- Name: products fk_products_brand; Type: FK CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.products
    ADD CONSTRAINT fk_products_brand FOREIGN KEY (brand_id) REFERENCES public.brands(id);


--
-- Name: products fk_products_category; Type: FK CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.products
    ADD CONSTRAINT fk_products_category FOREIGN KEY (category_id) REFERENCES public.categories(id);


--
-- Name: products fk_products_supplier; Type: FK CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.products
    ADD CONSTRAINT fk_products_supplier FOREIGN KEY (supplier_id) REFERENCES public.suppliers(id);


--
-- Name: DATABASE productsdb; Type: ACL; Schema: -; Owner: postgres
--

GRANT ALL ON DATABASE productsdb TO admin;


--
-- PostgreSQL database dump complete
--

