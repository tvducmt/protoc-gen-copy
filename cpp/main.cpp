#include <iostream>
#include <iomanip>
#include <cmath>
#include <cstdio>
#include <cstdlib>
#include <fstream>
using namespace std;


// TODO
int GetElement(float temperature, int option, int isRound) {
    float result;
 switch (option){
case 1: {
	result = (9/5) * temperature + 32;
		if (option == 1) // lam tron 2 chu so thap phan
			{
				result= (int)(result * 100 + .5); 
			}
		else
			{
				result= result;
			}
	return result;// #c to f
 }
 case 2: {
// f to c
	result = (5/9) * (temperature - 32);
			if (option == 1) // lam tron 2 chu so thap phan
			{
				result= (int)(result * 100 + .5); 
			}
		else
			{
				result= result;
			}
	return result; //c to f
 }
}
 
 return 0;

}

int main(){
    // float n = 12.12423
    // setiosflags(ios::n);
    // setprecison(n);
   cout << "sfgsf    d\n";
    cout << GetElement(30 ,2, 1)<<"\n";
    return 0;
}