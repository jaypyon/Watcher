# Watcher

Linux OS에서 System call을 이용하여 해쉬값에 해당하는 파일이 현재 사용자의 PC에 설치되어있는지 확인한다.

INPUT : hash.txt 
    hash.txt에 찾고자 하는 해쉬값(MD5, SHA-1) 리스트를 입력한다.

OUTPUT : result.txt
    [형식] 해쉬값의 종류 (MD5 or SHA-1) : 해쉬값 : 경로
    위와같은 형식으로 파일에 기록된다.
