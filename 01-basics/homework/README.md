# Homework

## 1. F.A.C.T.O.R.I.A.L.

The Gopher City is under attack! An evil mathematician has constructed a Frightening Autonomous, Calculating, Terrifying, Overwhelming, Ruthless, Intelligent, Adaptive, and Legendary robot, known as F.A.C.T.O.R.I.A.L. This colossal robot is on its way to destroy the city. Standing at a towering height of 120 meters, it is covered with numbers, making it extremely risky to approach.

Fortunately, there is a way to stop the robot: by entering a secret combination of numbers. This combination can be achieved by calculating the factorial of five numbers inscribed on the robot's body and then inputting the number of digits in each factorial result.

Our team of five brave superhero Gophers will confront the enemy. Each Gopher will shout a positive integer ranging from 1 to 10. Your task is to calculate the factorial of each shouted number and provide them with the number of digits in the resulting factorial. Armed with this information, they can enter the correct combination and destroy the robot.
Good luck! The fate of our city rests in your hands.

### Input:

5 random integers in the range [1,10] separated by space

### Output:

The number of digits of the factorial of each number, separated by space

### Example:

<table>
<thead>
  <tr>
    <th>Input</th>
    <th>Processing</th>
    <th>Output</th>
  </tr>
</thead>
<tbody>
  <tr>
  <td>1 2 3 4 5</td>
   <td>
    1! = 1 -> 1 digit<br>
    2! = 2 -> 1 digit<br>
    3! = 6 -> 1 digit<br>
    4! = 24 -> 2 digits<br>
    5! = 120 -> 3 digits<br>
    </td>
    <td>1 1 1 2 3</td>
  </tr>
    <tr>
  <td>6 6 1 4 7</td>
   <td>
    6! = 720 -> 3 digits<br>
    6! = 720 -> 3 digits<br>
    1! = 1 -> 1 digit<br>
    4! = 24 -> 2 digits<br>
    7! = 5040 -> 4 digits<br>
    </td>
    <td>3 3 1 2 4</td>
  </tr>
</tbody>
</table>

## 2. Club 34

Mayya Gopher is preparing for the party of a lifetime. She's dressed in her finest attire, sporting new shoes, and her hair and makeup are flawlessly done. While waiting in line outside the hottest dance club in town, Club 34, she overhears a rumor circulating through the crowd: there's a secret passcode to gain entry to the club.

Mayya observes as a few gophers ahead of her interact with the security. The guards utter "AB," and the first gopher responds with "21" and is granted entry. The next interaction unfolds: "Z!y X" - "262524"; followed by "Go" - "157". With her quick-wittedness, Mayya cracks the code: She takes the phrase, disregards the case and all punctuation, reverses it, and assigns numbers corresponding to the letters of the alphabet, starting from A as 1. While this approach works well for shorter phrases, the security guards soon begin using longer phrases.

To help Mayya and others join the party, a program is needed to convert any given phrase into a secret passcode. This program should follow the same rules: disregard case, reverse the phrase, and assign numbers based on the alphabet.

With this program, Mayya Gopher will be able to effortlessly obtain the secret passcode and join the vibrant party at Club 34!

### Input:

A phrase containing latin letters (lower- and uppercase), numbers, spaces and/or punctuation

### Output:

The corresponding position of the letters in the alphabet, where a is 1 and z is 26; in reverse order

### Example:

<table>
<thead>
  <tr>
    <th>Input</th>
    <th>Processing</th>
    <th>Output</th>
  </tr>
</thead>
<tbody>
<tr>
<td>AB</td>
<td>
A -> 1st letter<br>
B -> 2nd letter<br>
ab reverse is ba
ba corresponds to 21
</td>
<td>21</td>
  </tr>
  <tr>
<td>Z!y X</td>
<td>
Z -> 26th letter<br>
! -> not a letter, skipped<br>
y -> 25th letter<br>
space -> not a letter, skipped<br>
X -> 24th letter<br>
zyx reversed is xyz
xyz corresponds to 242526
</td>
<td>242526</td>
  </tr>
   <tr>
<td>Go</td>
<td>
G -> 7th letter<br>
o -> 15th letter<br>
go reversed is og
og corresponds to 157
</td>
<td>157</td>
  </tr>
</tbody>
</table>

## 3. Indiana Go-nes and the temple of Code

Indiana Go-nes stands in the temple of Code, fixated on an exceptionally precious artifact: a gopher crafted entirely from pure diamond. His desire to steal this treasure is undeniable, yet removing the statue from its pedestal would trigger a catastrophic event—a colossal boulder hurtling toward him, destined to flatten him to oblivion. Fortunately, our adventurer possesses a bag filled with pebbles, a suitable substitute for the statue.

Indie is keenly aware that the temple's developers are meticulous in their design; even the slightest gram disparity between the statue and the substituting pebbles could spell disaster. Through careful inspection, Indie notices that the statue is constructed from numerous smaller crystal pieces, each denoted with a number signifying its weight in grams.

Now, it's up to you to assist our brave protagonist in successfully accomplishing his mission. He will provide a sequence of numbers—both positive and negative. The positive values represent the weight of the statue shards, while the negative values represent the weight of the pebbles. As the stream concludes, he will relay the value 0. Upon receiving this signal, your task is to ascertain whether Indie can seamlessly substitute the statue with the pebbles. If the combined weight of the pebbles equals that of the statue, signify this triumph with a resounding "Yes" Otherwise, inform Indie of the precise weight difference between the two objects.

### Input:

A stream of positive and negative integers, ending with 0, separated by space

### Output:

"Yes" if the absolute sum of the negative equals the sum of the positive. Otherwise - the absolute difference between the sums of the two sequences

### Example:

<table>
<thead>
  <tr>
    <th>Input</th>
    <th>Processing</th>
    <th>Output</th>
  </tr>
</thead>
<tbody>
<tr>
<td>1 -1 12 1 1 1 -5 3 -7 -10 4 0</td>
<td>
Sum of positive: 1 + 12 + 1 + 1 + 1 + 3 + 4 = 23 <br>
Sum of negative: -1 + -5 + -7 + -10 = -23
|-23| = 23 => therefore they are equal
</td>
<td>"Yes"</td>
  </tr>
  <tr>
<td>5 -4 12 -7 -3 -2 0</td>
<td>
Sum of positive: 5 + 12 = 17 <br>
Sum of negative: -4 + -7 + -3 + -2 = -16
|-16| != 17 => therefore they are not equal
17 - 16 = 1
</td>
<td>1</td>
  </tr>
    <tr>
<td>10 18 30 0</td>
<td>
Sum of positive: 10 + 18 + 30 = 58 <br>
No negatives
0 != 58 => therefore they are not equal
</td>
<td>58</td>
  </tr>
  </tbody>
</table>

## 4. Raffle

The Sveti Gopher has meticulously prepared a delightful assortment of gifts for the enthusiastic participants attending the highly anticipated Golang Workshop. However, a dilemma arises when it becomes apparent that the number of prizes is smaller than the number of eager participants. To ensure an fair distribution, a raffle is proposed. Each participant will draw a ticket with a random number, and the winning tickets shall be those that possess a minimum of two identical digits.

In the post-presentation frenzy of answering questions, Sveti finds herself immersed in the commotion and seeks your assistance in determining the prized tickets. Additionally, to speed up verifying a winning ticket, Sveti needs to know the last repeating digit. Your mission is to write a program that accepts a natural number as input and returns the last repeating digit within it. If no repeating digit is discovered, the program should return the word "No".

### Input:

A single line containing a natural number

### Output:

The program should output the last repeating digit in the number, if there are any repeated digits in the representation of n. Otherwise, the program should output the word "No".

### Examples:

<table>
<thead>
  <tr>
    <th>Input</th>
    <th>Output</th>
  </tr>
</thead>
<tbody>
  <tr>
    <td>234256</td>
    <td>2</td>
  </tr>
   <tr>
    <td>676678</td>
    <td>7</td>
  </tr>
   <tr>
    <td>2345</td>
    <td>No</td>
  </tr>
</tbody>
</table>

## 5. Censor

Oh no! The client you're working for has been hacked, and their files have fallen victim to a mischievous virus. The virus has gone wild, replacing random letters with sneaky asterisks. It's like a secret code, but without the fun!

The client, in a state of panic, politely asks if you could stay a little later to help them recover their files. However, you have an exciting date planned and don't want to stay too late. Time is of the essence!

Your mission, should you choose to accept it, is to implement an "uncensor" function that swiftly reveals the hidden letters and restores the client's files. The faster you can uncensor, the sooner you can head off to your date.

### Input:

Two strings each on a new line: the first one containing the message with asterisks and the second one containing the missing letters

### Output:

The uncensored message

### Examples:

<table>
<thead>
  <tr>
    <th>Input</th>
    <th>Output</th>
  </tr>
</thead>
<tbody>
  <tr>
    <td>*h*s *s v*ry *tr*ng*<br>Tiiesae</td>
    <td>2</td>
  </tr>
   <tr>
    <td>A**Z*N*<br>MAIG</td>
    <td>AMAZING</td>
  </tr>
   <tr>
    <td>xyz<br></td>
    <td>xyz</td>
  </tr>
</tbody>
</table>

### Notes:

- Expect the discovered letters to be given in the correct order.
- The number of discovered letters will match the number of censored ones.
- Any character can be censored.

## 6. Genie

Once upon a time, in a land of quirky gadgets, there existed a lamp that possessed an amusingly unpredictable personality. This lamp had a knack for changing its color and state, keeping everyone on their toes. It would flicker on and off, displaying a vivid spectrum of hues ranging from red, green, orange, yellow, blue, to purple.

Within this mesmerizing scenery, Gopher Niki stumbled upon the mystical lamp, stirred by rumors that a wondrous genie would materialize to grant three wishes only when the lamp radiated a brilliant red glow while being in an activated state.

Help our fortunate gopher in unlocking the secrets of the lamp and he will give you one of his wishes in return. He will give you a base struct that implements a Lamp, and your task is to extend the Lamp struct in Go by adding the following methods:

- `NewLamp(color string, state string)`: A method that serves as the magical lamp's constructor, creating a new instance with the provided color and state values. The lamp by default has the color black and is off.
- `GetColor() string`: A method that reveals the current color of the lamp.
- `SetColor(color string)`: A method that allows you to play fashion designer for the lamp by setting its color to the specified hue only in the range of red, green, orange, yellow, blue, to purple. If the given color is not in this range: color the lamp black as a warning.
- `GetState() string`: A method that unveils the lamp's current state, be it on or off.
- `SetState(state string)`: A method that gives you the power to control the lamp's whimsical state, turning it on or off at will.
- `SummonGenie() bool:` return weather the genie can be summoned

```go

type Lamp struct {
    color string
    state string
}

// Your code
```

### Input

Six commands for the lamp in total: 3 colors (one word) and 3 states (on or off). The order of colors and states can be mixed

### Output

The state and the color of the lamp after each new input and at the end of the commands: if the genie can be summoned or not

### Examples:

<table>
<thead>
  <tr>
    <th>Input</th>
    <th>Output</th>
  </tr>
</thead>
<tbody>
  <tr>
    <td>blue on purple off red on</td>
    <td>
    blue off<br>
    blue on<br>
    purple on<br>
    purple off<br>
    red off<br>
    red on<br>
    Yes
    </td>
  </tr>
   <tr>
    <td>orange off red off turquoise on</td>
    <td>
    orange off<br>
    orange off<br>
    red off<br>
    red off<br>
    black off<br>
    black on<br>
    No
    </td>
  </tr>
   <tr>
    <td>yellow purple on green off on</td>
    <td>
    yellow off<br>
    purple off<br>
    purple on<br>
    green on<br>
    green off<br>
    green on<br>
    No
    </td>
  </tr>
</tbody>
</table>
