thoughts:

lex into tokens
parser analyzes the tokens
certain strings of tokens create different AST nodes
recurse over AST until all checks are finished
final recurse is a translation of the AST into C++
