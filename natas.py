from solutions.natas0 import solve0
from solutions.natas1 import solve1
from solutions.natas2 import solve2
from solutions.natas3 import solve3
from solutions.natas4 import solve4
from solutions.natas5 import solve5
from solutions.natas6 import solve6
from solutions.natas7 import solve7
from solutions.natas8 import solve8
from solutions.natas9 import solve9
from solutions.natas10 import solve10
from solutions.natas11 import solve11
from solutions.natas12 import solve12
from solutions.natas13 import solve13
from solutions.natas14 import solve14
from solutions.natas15 import solve15
from solutions.natas16 import solve16
from solutions.natas17 import solve17
from solutions.natas18 import solve18
from solutions.natas19 import solve19
from solutions.natas20 import solve20
from solutions.natas21 import solve21
from solutions.natas22 import solve22
from solutions.natas23 import solve23
from solutions.natas24 import solve24
from solutions.natas25 import solve25
from solutions.natas26 import solve26
from solutions.natas27 import solve27



if __name__ == '__main__':
    natasPWs = ["natas0"]
    print("{:<8} | {:<10} | {:<7} | {:<}".format("Puzzle", "Solved", "Username", "Password"))
    print("="*66)
    for i in range(28):
        print("{:<8} | {:<10} | {:<7} | {:<}".format("Natas " + str(i), "Not Solved", "", ""), end="")
        eval('natasPWs.append(solve'+str(i)+'("natas" + str(len(natasPWs)-1), natasPWs[-1]))')
        print("\r{:<8} | {:<10} | {:<7} | {:<}".format("Natas " + str(i), "Solved", "natas" + str(i), natasPWs[-1]))

