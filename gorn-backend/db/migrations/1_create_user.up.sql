CREATE TABLE public."user"
(
    id serial NOT NULL,
    firstname  text NOT NULL,
    lastname  text NOT NULL,
    PRIMARY KEY (id)
);

CREATE TABLE public."order"
(
    id serial NOT NULL,
    firstname  text NOT NULL,
    lastname  text NOT NULL,
    phone  text NOT NULL,
    city  text NOT NULL,
    postalcode text NOT NULL,
    state text NOT NULL,
    country text NOT NULL,
    crayon00 text NOT NULL,
    crayon01 text NOT NULL,
    crayon02 text NOT NULL,
    crayon03 text NOT NULL,
    crayon04 text NOT NULL,
    crayon05 text NOT NULL,
    crayon06 text NOT NULL,
    crayon07 text NOT NULL,
    crayon08 text NOT NULL,
    crayon09 text NOT NULL,
    crayon10 text NOT NULL,
    crayon11 text NOT NULL,
    crayon12 text NOT NULL,
    crayon13 text NOT NULL,
    crayon14 text NOT NULL,
    crayon15 text NOT NULL,
    crayon16 text NOT NULL,
    crayon17 text NOT NULL,
    crayon18 text NOT NULL,
    crayon19 text NOT NULL,
    PRIMARY KEY (id)
);

ALTER TABLE public."user"
    OWNER to pgndsusa;

ALTER TABLE public."order"
    OWNER to pgndsusa;