Modul yg menyelesaikan masalah

The alphabet, from 'a' to 'z', is assigned weights corresponding to their ordinal positions. For instance, 'a' has a weight of 1, 'b' has a weight of 2, and so on, with 'z' having a weight of 26. The weight of a string is calculated by summing the weights of all its characters. Consider the string "gits" > gits = 7 + 9 + 20 + 19 = 55.

Example:
Input:
Given a string "abbcccd" and an array of queries "[1, 3, 9, 8]". The queries are used to determine the status of the numbers in the queries based on the following rules:
1. If a number in the queries is equal to the weight of a character or substring, return "Yes".
2. If a number in the queries is different from the weight of a character or substring, return "No".
a = 1
b = 2
bb = 4
c = 3
cc = 6
ccc = 9
d = 4 
Output: [Yes, Yes, Yes, No]
Explanation: 
1. 1=> Yes, 3 => Yes, 9 => Yes, and 8 => No.
2. Based on the character and substring weighting of "abbcccd", the status of the queries "[1, 3, 9, 8]" is "[Yes, Yes, Yes, No]".

Rule:
For repeated and consecutive characters, the weighting should be specified for each substring from the first occurrence to the n-th occurrence. Example: bbccc => b, bb, c, cc, ccc.

Task: 
Create a function to solve the Weighted Strings problem!

Determining whether a string of brackets is balanced is a fundamental problem in computer science, particularly in parsing and validation of expressions. A balanced string of brackets consists of pairs of opening and closing brackets, where each opening bracket has a corresponding closing bracket of the same type, and the brackets are nested properly.

Given a string of brackets, determine whether it is balanced. The string may contain different types of brackets, including parentheses '()', square brackets '[]', and curly braces '{}'. The brackets may be separated by whitespace or not.

Example 1:
Input: { [ ( ) ] }
Output: YES
Explanation: Each bracket is balanced, with every opening bracket having a corresponding closing bracket of the same type.
opening : { }
opening : [ ]
opening : ( }

Example 2:
Input: { [ ( ] ) }
Output: NO
Explanation: The string { [ ( ] ) } is not balanced for the characters enclosed by the curly braces '{}', namely '[ ( ]'.

Example 3:
Input: { ( ( [ ] ) [ ] ) [ ] }
Output: YES
Explanation: Each bracket is balanced, with every opening bracket having a corresponding closing bracket of the same type, even though the bracket structure is irregular.

Rules:
1. Allowed bracket/character types: ( , ) , { , } , or [ , ]. 
2. Brackets can be separated with or without whitespace.
3. Check for matching pairs of opening and closing brackets and return a string "YES" or "NO" accordingly.

Task:
1. Create a function to find Balanced Brackets with the Lowest Complexity!
2. Determine the complexity of your code. Explain the details of the complexity analysis for answer No. 2 and include it in the README of your repository!

Given a string representing a number, the goal is to find the highest possible palindrome that can be formed by changing at most 'k' digits in the string. A palindrome is a number that reads the same backward as forward. 

You are given a string 's' representing a number and an integer 'k'. The task is to find the highest palindrome that can be formed by changing at most 'k' digits in the string 's'.


Example 1:
Input:
s: 3943 
k: 1 
Palindrome:
1. 3943  => 3993
2. 3943 => 3443
Output: 3993
Explanation: Among the possible palindromes obtained, the highest palindrome is 3993 since 3993 > 3443.

Example 2:
Input:
s: 932239 
k: 2
Palindrome: 
1. 932239 => palindrome 
2. Further Replacement (k = 2): => 992299
Output: 992299 
Explanation: Among the possible palindromes obtained, the highest palindrome is 992299 since 992299 > 932239.

Rules:
1. If a palindrome cannot be formed from the given string even after 'k' replacements, and the string does not represent a valid number, output '-1'.
2. Do not use built-in functions or tools for searching, filtering, or sorting.
3. Avoid using loops.
4. Only use recursion to solve the problem.

Task:
Create a recursive function to solve the Highest Palindrome problem!
