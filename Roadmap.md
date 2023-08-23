# Roadmap

## Protobuf AST
![ast](./ast.png)

## Message Features
- [ ] oneof: 어떻게?
- [x] map: Json Editor?
- [x] scala: 주 과제
- [ ] repeated: Scala, messsage의 list형, 어떻게?
- [x] optional
- [x] expose
- [x] message in message

## Scala Field Features
- [ ] select with external source: 툴파트랑 필드 이름 논의해야 함
- [x] string options: min, max, regex?
- [x] number options: ?
- [x] boolean type: checkbox or toggle
- [x] enum type: select or radio set

## Scala Type
1. String
2. Number
3. Enum
4. Boolean

## Well-known type
- [ ] Any
- [ ] timestamp
- [ ] duration
- [ ] k8s.io.apimachinery.pkg.apis.util.v1.IntOrString
- [ ] k8s.io.api.pkg.core.v1.Volume
- [x] map
- [x] boolean-repeated
- [ ] string repeated (aka. multi string)
- https://protobuf.dev/reference/protobuf/google.protobuf/

## 추후 과제
- [ ] 2차원 Location Override
- [x] Property Override
- [x] protobuf package name change to [pubg.venus]
