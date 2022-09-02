# golang-test-dihardja
golang technical test dihardja


questions :


Logical Programming Test (20 mins)

1. A palindrome is a word, phrase, number, or other sequence of characters which reads the same backward or forward.(Wikipedia). Create a function to check wether the word is palindrome or not


Given a string A , print Yes if it is a palindrome, print No otherwise.
Constraints
	 will consist of at most 10 lower case english letters.
Sample Input
     madam			

Sample Output
Yes		
	
2. Berikan Fungsi untuk print segitiga berikut:
	*
           **
          ***
         ****
        *****

	
3. Given a number N. Find the minimum number of operations required to reach N starting from 0. You have 2 operations available: 
Double the number
Add one to number

	Sample Input
		N = 8


	Sample Output
		4

	Explanation
		0 + 1 = 1, 1 + 1 = 2, 2 * 2 = 4, 4 * 2 = 8

4. Given a square matrix, calculate the absolute difference between the sums of its diagonals.
For example, the square matrix  is shown below:
1 2 3
4 5 6
9 8 9  
 
The left-to-right diagonal = 1+5+9 = 15 The right to left diagonal = 3+5+9 = 17 Their absolute difference is |15-17| = 2
Function description
diagonalDifference takes the following parameter:
int arr[n][m]: an array of integers
Return
int: the absolute diagonal difference
Input Format
The first line contains a single integer, , the number of rows and columns in the square matrix 
Each of the next  lines describes a row, , and consists of  space-separated integers .
Constraints 
-100 <=  arr[i][j] <= 100
 
 
Output Format
Return the absolute difference between the sums of the matrix's two diagonals as a single integer.
Sample Input
3
11 2 4
4 5 6
10 8 -12
 
Sample Output
15
 
Explanation
The primary diagonal is:
11
    5
      -12
 
Sum across the primary diagonal: 11 + 5 - 12 = 4
The secondary diagonal is:
          4
      5
10
 
Sum across the secondary diagonal: 4 + 5 + 10 = 19
Difference: |4 - 19| = 15

