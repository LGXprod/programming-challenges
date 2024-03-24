def cons(a, b):
    def pair(f):
        return f(a, b)

    return pair

def car(pair):
  def getFirst(a, _):
    return a
  
  return pair(getFirst)

def cdr(pair):
  def getLast(_, b):
    return b
  
  return pair(getLast)

pair = cons(3, 4)

print("car of pair:", car(pair))
print("cdr of pair:", cdr(pair))


pair = cons(6, 2)

print("car of pair:", car(pair))
print("cdr of pair:", cdr(pair))
