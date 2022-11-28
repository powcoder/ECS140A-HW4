https://powcoder.com
代写代考加微信 powcoder
Assignment Project Exam Help
Add WeChat powcoder
https://powcoder.com
代写代考加微信 powcoder
Assignment Project Exam Help
Add WeChat powcoder
% A list is a 1-D array of numbers.
% A matrix is a 2-D array of numbers, stored in row-major order.

% You may define helper functions here.
match_in_order([H|[TH|_]], A, B) :- H = A, TH = B.

at_least_one_match(List, A, B) :- 
    match_in_order(List, A, B); match_in_order(List, B, A).

% loop_list([], _, _) :- fail.
loop_list([H|T], A, B) :- 
    at_least_one_match([H|T], A, B); loop_list(T, A, B).

% are_adjacent(List, A, B) returns true iff A and B are neighbors in List.
are_adjacent(List, A, B) :-
   loop_list(List, A, B).

% transpose(Matrix, Answer) returns true iff Answer is the transpose of the 2D
% matrix Matrix
transpose([], []).
transpose([MH|MT], Answer) :-
    transpose(MH, [MH|MT], Answer).

transpose([], _, []).
transpose([_|MT], Matrix, [AH|AT]) :-
    extract_first_col(Matrix, AH, RestMatrix),
    transpose(MT, RestMatrix, AT).

extract_first_col([], [], []).
extract_first_col([[H|T]|RestRows], [H|TT], [T|RestT]) :-
    extract_first_col(RestRows, TT, RestT).

% are_neighbors(Matrix, A, B) returns true iff A and B are neighbors in the 2D
% matrix Matrix.
% are_neighbors([], _, _) :- fail.
are_neighbors(Matrix, A, B) :-
    loop_matrix(Matrix, A, B);
    check_transpose(Matrix, A, B).

check_transpose(Matrix, A, B) :-
    transpose(Matrix, Transpose),
    loop_matrix(Transpose, A, B).

loop_matrix([H|T], A, B) :-
    are_adjacent(H, A, B);
    loop_matrix(T, A, B).

