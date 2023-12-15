CryptoFungi Project

# 크립토풍기 프로젝트 실행 방법

## 1. BlockChain Network 실행

./startNetwork.sh

## 2. 체인코드 배포
./deployCC.sh
### 위 sh 실행 시 아래의 CC 배포
#### 버섯공장(fungusfactory)

#### 먹이공장(feedfactory)

## 3. 체인코드 테스트 ( Initalize를 위해 반드시 수행필요 )
./testCC.sh

* 카우치DB 확인 url : http://localhost:5984/_utils/
* id : admin
* password : adminpw

## 4. Application 실행
./startApplication.sh

## 5. Web Client 접속
웹브라우저 : localhost:3000

## END : 네트워크 종료
./networkDown.sh

# 업데이트 강의자료와 code는 ftp://202.30.3.201 와 본 github로 공유 예정입니다.
# 하이퍼레저 패브릭을 공부하시면서 도움이 필요하신 경우 gcloud@ajou.ac.kr로 이메일 보내시면 도와드리겠습니다.
##그럼 좋은 하루 되세요
