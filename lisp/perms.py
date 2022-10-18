
def perms(s):
    if len(s) == 2:
        return [s, s[::-1]]
    result = []
    for i, c in enumerate(s):
        substr = s[:i] + s[i+1:]
        for subperm in perms(substr):
            result.append(c+subperm)
    return result

print(perms('abcd'))
