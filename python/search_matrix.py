"""Given a matrix that is organized such that the numbers will always be sorted left to right, and the first number of each row will always be greater than the last element of the last row (mat[i][0] > mat[i - 1][-1]), search for a specific value in the matrix and return whether it exists."""

def find_row(mat, val):
  found_row = False
  first, last = 0, len(mat) - 1
  while not found_row and first <= last:
      row = first + (last-first)//2
      found_row = val >= mat[row][0] and val <= mat[row][-1]
      if found_row:
        return mat[row]
      elif val < mat[row][0]:
        last = row - 1
      else:
        first = row + 1
  return []

def find_val(row, val):
  first, last = 0, len(row) - 1
  while first <= last:
    i = first + (last-first)//2
    if row[i] == val:
      return True
    elif val < row[i]:
      last = i - 1
    else:
      first = i + 1
  return False

def searchMatrix(mat, value):
  row = find_row(mat, value)
  if row:
    return find_val(row, value)
  return False

  
mat = [
  [1, 3, 5, 8],
  [10, 11, 15, 16],
  [24, 27, 30, 31],
]

print(searchMatrix(mat, 4))
# False

print(searchMatrix(mat, 10))