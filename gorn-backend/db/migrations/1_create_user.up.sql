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

INSERT INTO public."order" (firstname, lastname, phone, city, postalcode, state, country, crayon00, crayon01, crayon02, crayon03, crayon04, crayon05, crayon06, crayon07, crayon08, crayon09, crayon10, crayon11, crayon12, crayon13, crayon14, crayon15, crayon16, crayon17, crayon18, crayon19)
VALUES ('firstname00','lastname00','+1 555 555 5555','city00','72000','ST','CO','#000000','#FFFFFF','#000000','#FFFFFF','#000000','#FFFFFF','#000000','#FFFFFF','#000000','#FFFFFF','#000000','#FFFFFF','#000000','#FFFFFF','#000000','#FFFFFF','#000000','#FFFFFF','#000000','#FFFFFF')
