https://powcoder.com
代写代考加微信 powcoder
Assignment Project Exam Help
Add WeChat powcoder
https://powcoder.com
代写代考加微信 powcoder
Assignment Project Exam Help
Add WeChat powcoder
/* All novels published either during the year 1953 or during the year 1996*/
year_1953_1996_novels(Book) :- 
    novel(Book, 1953); novel(Book, 1996).

/* List of all novels published during the period 1800 to 1900 (not inclusive)*/
period_1800_1900_novels(Book) :-
    novel(Book, Year), within_1800_1900(Year).

within_1800_1900(Year) :- Year >= 1800, Year =< 1900.

/* Characters who are fans of LOTR */
lotr_fans(Fan) :-
    fan(Fan, Books), has_lotr(Books).

% has_lotr([]) :- fail.
has_lotr([H|T]) :- H = the_lord_of_the_rings; has_lotr(T).

/* Authors of the novels that heckles is fan of. */
heckles_idols(Author) :-
    fan(heckles, Books), loop_books(Books, Author).

% loop_books([], _) :- fail.
loop_books([H|T], Author) :- 
    write_by_author(H, Author); loop_books(T, Author).

write_by_author(Book, Author) :- 
    author(Author, Books), in_book_list(Books, Book).

% in_book_list([], _) :- fail.
in_book_list([H|T], Book) :- 
    H = Book; in_book_list(T, Book).

/* Characters who are fans of any of Robert Heinlein's novels */
heinlein_fans(Fan) :-
    fan(Fan, Books), loop_books(Books, robert_heinlein).

/* Novels common between either of Phoebe, Ross, and Monica */
mutual_novels(Book) :-
    common_between(phoebe, ross, monica, Book).

common_between(A, B, Book) :- 
    in_fan_books(A, Book), in_fan_books(B, Book).

common_between(A, B, C, Book) :- 
    common_between(A, B, Book); 
    common_between(A, C, Book);
    common_between(B, C, Book).

in_fan_books(Fan, Book) :- 
    fan(Fan, Books), in_book_list(Books, Book).
