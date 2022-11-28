https://powcoder.com
代写代考加微信 powcoder
Assignment Project Exam Help
Add WeChat powcoder
https://powcoder.com
代写代考加微信 powcoder
Assignment Project Exam Help
Add WeChat powcoder
reachable(_, StartState, FinalState, []) :- StartState = FinalState.
reachable(Nfa, StartState, FinalState, [Input|T]) :-
    transition(Nfa, StartState, Input, NextStates),
    loop_states(Nfa, NextStates, T, FinalState).

% loop_states(_, [], _, _) :- fail.
loop_states(Nfa, [NextState|T], Input, FinalState) :- 
    reachable(Nfa, NextState, FinalState, Input), !;
    loop_states(Nfa, T, Input, FinalState).
