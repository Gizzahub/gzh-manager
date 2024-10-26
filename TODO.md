# TODO to setclone

## 중요!!
테스트 할 때 Gizzahub에다 하지 않기 ^^
동작 잘 하는거 확인했다 ^^b al

## Setclone
- [ ] setclone cli
  - [x] -o organization name: `-o ScriptonBasestar`
  - [x] -t targetPath: `-t $HOME/mywork/ScriptonBasestar`
  - [x] -p proto: `-p https` or `-p ssh`
  - [x] -a auth: `-a token` 토큰이 있으면 private repo도 clone
  - [ ] -s strategy: `-s reset` or `-s pull` or `-s fetch`
- [ ] setclone.yaml 설정파일 지원
- [ ] setclone.yaml override setclone.home.yaml setclone.work.yaml. kustomize 이용
- [ ] setclone.yaml 설정파일 예시
- [ ] setclone.yaml schema
- 
### Github Org
- [x] gituhb org setclone cli

### Gitlab Org
- [ ] gitlab group setclone cli

### Ssh config
- [ ] ssh config 설정
- [ ] config 설정 만들어내기

## Always Latest
- [ ] asdf
- [ ] brew
- [ ] sdkman
- [ ] port
- [ ] apt-get
- [ ] rbenv

업데이트 방식
- minor latest
- major latest

## 개발환경 Save & Load
- [ ] kubeconfig
- [ ] docker config
- [ ] aws config
- [ ] aws credentials
- [ ] gcloud config
- [ ] gcloud credentials
- [ ] ssh config
- ...

## 네트워크 옮길 때
- [ ] daemon 모니터링
- [ ] hook wifi change event? -> action
- [ ] action: vpn, dns, proxy, host 등 변경
