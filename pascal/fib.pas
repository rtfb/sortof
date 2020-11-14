program fibonacci;

var n: integer;

function fib(n: integer): integer;
begin
	if n = 1 then
		fib := 1
	else if n = 2 then
		fib := 1
	else
		fib := fib(n-1) + fib(n-2)
end;

begin
	for n := 1 to 11 do begin
		WriteLn(n, ': ', fib(n))
	end;
end.
