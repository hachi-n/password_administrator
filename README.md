# password_administrator
 
## Dependency 
* sqlite3 

## config/application.ini 
- web section is dummy. 
    - I will make it soon
## cli 
please command -h 

## example 
```
./password_administrator -s ${serviceName} -c ${categoryName}
************************* New Password *************************
Service:  serviceName
Category: categoryName
Password: M2<dSxsoBhQ6A"cJ
```

### interactive mode
```
./password_administrator
```

### password Strength
```
./password_administrator -l 100 -n 5 -m 6
************************* New Password *************************
Service:  hogehoge
Category: private
Password: ozkew1DjzcqLJdJxaRZ:FYFchERDP&PvznamnfjNyPlN7klOxtegprJDHXTbyvzO2+vTVKox[VRy4BTNnEWvHQPeTOm!ym)2npyz
```
l ... length 
n ... digit number 
m ... symbol number 
default value  is config/application.ini section 


### password List 
```
./password_administrator -L
************************* 1 *************************
Service:  adfsfadfdsaf
Category: private
Password: S}3vr{1CuJxZyQZf

************************* 2 *************************
Service:  serviceName
Category: categoryName
Password: M2<dSxsoBhQ6A"cJ

************************* 3 *************************
Service:  hoge
Category: private
Password: rK|qW/WNBy5l8lNJ

************************* 4 *************************
Service:  hogehoge
Category: private
Password: ozkew1DjzcqLJdJxaRZ:FYFchERDP&PvznamnfjNyPlN7klOxtegprJDHXTbyvzO2+vTVKox[VRy4BTNnEWvHQPeTOm!ym)2npyz

```
