# Zadanie 1

## Część obowiązkowa

### 1.

Wykorzystuję algorytm obliczania n-tej liczby Fibonacciego za pomocą potęgowania macierzy. Korzystam ze wzoru:<br />
![wzor](https://github.com/Adrian3753/zad1/blob/main/screeny/screen2.png)<br /> 
Ze wzoru można odczytać, że podniesienie lewej macierzy do potęgi n-tej, da nam macierz, której elementy są liczbami Fibonacciego. 
Wystarczy, że podniosę macierz do potęgi (n-1) - wtedy porządanym wynikiem będzie element w pierwszej kolumnie i pierwszym wierszu macierzy. 
Złożność algorytmu to O(log(n-1)).<br /><br />
Wpisuję ``gh auth login --with-token < ~/.ssh/git_token.txt``, aby zalogować się na moim koncie na GitHub.<br />
Inicjalizuje lokalne repozytorium git ``git init -b main``<br />
Dodaję do utworzonego repozytorium plik fibCalc.go i tworzę pierwszy commit ``git add . && git commit -m "Dodajemy fibCalc.go"``<br />
Tworzę nowe repozytorium na GitHub i kopiuję tam zainicjalizowane repozytorium git ``gh repo create zad1 --public --source=. --remote=zad1 --push``
![komendy](https://github.com/Adrian3753/zad1/blob/main/screeny/screen1.png)<br />

### 2.
Wpisuję ``docker build --tag fibcalc .``, aby zbudować obraz.<br />
![komendy](https://github.com/Adrian3753/zad1/blob/main/screeny/screen3.png)<br />
Wpisuję ``docker run -it --rm fibcalc``, aby uruchomić kontener. <br />
Poniższy screen przedstawia testowanie aplikacji. Sprawdziłem czy działa walidacja, a następnie czy jest zwracany poprawny wynik dla przykładowych dwóch wartości.<br />
![komendy](https://github.com/Adrian3753/zad1/blob/main/screeny/screen4.png)<br />

### 3.
Nazewnictwo obrazów zgodnie z metodą semver. Wersje ozaczam w postaci x.y.z, gdzie zmieniam:<br >
"x", gdy dokonuję zmian niekompatybilnych z API,<br />
"y", gdy dodaję nową funkcjonalność, która jest kompatybilna z poprzednimi wersjami,<br />
"z", gdy naprawiam błąd, nie wpływaąc na kompatybilność z poprzednimi wersjami.<br /><br />

W celu zastosowania nazewnictwa obrazów po dodaniu zmian w repozytorium i stworzeniu commita ``git add . && git commit -m "Proba progresu nr 3"``, dodałem tag ``git tag 1.0.0``, wysłałem zmiany do zdalengo repozytorium ```git push```, a następnie wysłałem utworzony tag do zdalnego repozytorium ```git push --tags```
![komendy](https://github.com/Adrian3753/zad1/blob/main/screeny/screen5.png)<br />
