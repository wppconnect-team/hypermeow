# Paridade isolada com a Zuck v2.4.2 na v0.1.40

A release de referência é a [v0.1.40 do server-go](https://github.com/wppconnect-team/wppconnect-server-go/releases/tag/v0.1.40), commit `0b4b3091377a42280145df134d1625580f00312c`. O manifesto OCI usa `sha256:b9a33faca0a1c8cfef7f6eab7b6d33324f64b56c04acefc52e5cc3cda401abe3` e inclui `linux/amd64` e `linux/arm64`. As verificações desta rodada rodaram com contas, dados e destinos sintéticos em processos e contêineres descartáveis.

## Comparação de contratos

- [x] A matriz congelada tem 303 operações da Zuck. O OpenAPI da v0.1.40 tem 170 operações e o Swagger do Node fixado em `37b7bc1a662b5262b8276597701f38c4dbcad58a` tem 158. A interseção literal por método e caminho é 87, com 83 operações exclusivas do Go e 71 do Node. A contagem descreve a superfície HTTP, incluindo métodos experimentais ou indisponíveis declarados no contrato.
- [x] Sete funções passaram para os 22 contratos nativos mapeados. Os dez diferenciais antes identificados entre o Node e o Go têm contratos Go testados. [Matriz](parity-contract-matrix.md) · [Mapeamento](parity-progress.json)
- [x] Z290, Z271, Z282, Z285 e Z283 passaram nos testes de exclusão de canal e coleção, criação, edição e exclusão de produtos. As fixtures exercitam ID e papel exigidos para excluir canal, limites de lote, precisão de preço, moeda, revisão do produto, visibilidade, contagem parcial e recuperação idempotente.
- [x] Um processo Go desconectado foi encerrado e recriado com o mesmo SQLite. Os cinco resultados HTTP 404 originais continuaram consultáveis. A repetição devolveu o mesmo corpo e o cabeçalho de replay. A mesma chave com outro corpo recebeu 409. [Teste](check-isolated-mutation-replay.py)
- [ ] Executar a comparação HTTP com a Zuck v2.4.2 licenciada. O Swagger permite comparar campos e códigos declarados. Os corpos reais, erros operacionais, efeitos remotos, recuperações e limites exigem executar ambos os servidores. A referência encerrou a inicialização nos ensaios anteriores sem `LICENSE_KEY`. [Registro da referência](reference-runtime-acceptance.json)

## Sessões e restauração

- [x] Um contêiner descartável da imagem publicada verificou `pair-phone` sem credencial (401), com token de outra sessão (401) e com telefone inválido (400). O logout desconectado retornou 409. `close-session` retornou 200 e o estado `CLOSED` permaneceu após reiniciar.
- [x] As fixtures do Manager testaram validação, cancelamento e geração de código pelo cliente simulado. Testaram também preservação de credenciais quando o logout remoto falha e revogação local quando tem sucesso. A migração SQLite da v16 à v21 verificou chaves, segredo WASA, registro histórico e repetição do upgrade. [Ciclo de sessão](https://github.com/wppconnect-team/wppconnect-server-go/blob/0b4b3091377a42280145df134d1625580f00312c/internal/session/lifecycle_test.go) · [Pareamento](https://github.com/wppconnect-team/wppconnect-server-go/blob/0b4b3091377a42280145df134d1625580f00312c/internal/session/settings_test.go)
- [ ] Confirmar vínculo por código e revogação remota em conta descartável. O endpoint de código aguarda a prontidão do transporte inicial. Nenhum código foi solicitado a um telefone real nesta rodada.

## Aceites pendentes

| Dependência | Critério de encerramento |
| --- | --- |
| WhatsApp Business com catálogo de teste | Ler e paginar produtos e coleções, validar efeitos de criação, edição, exclusão e visibilidade, conferir revisão, leitura no comprador e estado do carrinho. |
| Instalação licenciada da Zuck v2.4.2 | Executar a matriz comportamental lado a lado. Validar corpos, falhas, repetição, isolamento, integração e efeitos remotos. |
| Decisões de teste já registradas | Exclusão real de canal segue restrita a testes isolados. Segunda conta e clientes Android, iPhone e Web continuam adiados. Monitoramento externo continua pendente de serviço e destino. |
| Implementação e testes restantes | Anexos bidirecionais no Chatwoot, arquivamento automático com opt-in, pedidos e pagamentos, entrega de Flows e interativos, áudio/vídeo e chamadas em grupo, carga de protocolo com metas. |

A confirmação de texto do Chatwoot e o upload explícito com retenção em S3 cobrem partes das duas primeiras integrações. Os critérios completos para lançamentos seguintes estão em [próximas entregas](parity-remaining.md).
