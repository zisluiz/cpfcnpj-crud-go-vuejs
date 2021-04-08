//db = new Mongo().getDB("cpfcnpj");

let res = [
  db.document.drop(),
  db.createCollection('document'),
  db.document.createIndex({ identityNumber: 1 }, { unique: true }),
  db.document.createIndex({ name: 1 }),
  db.document.insert({ name: 'Luiz', identityNumber: '31717241620', blocked: false }),
  db.document.insert({ name: 'Paulo', identityNumber: '41637711581', blocked: false }),
  db.document.insert({ name: 'João', identityNumber: '24455286403', blocked: true }),
  db.document.insert({ name: 'Maria', identityNumber: '54096788260', blocked: false }),
  db.document.insert({ name: 'Paula', identityNumber: '06016982318', blocked: false }),
  db.document.insert({ name: 'Empresa X LTDA', identityNumber: '66048673000120', blocked: false }),
]

printjson(res)