# Progresso sobre a matriz Zuck v2.4.2

A [matriz de 303 operações](parity-contract-matrix.md) é o inventário da versão comparada. Esta tabela registra o que mudou depois da medição, sem alterar o baseline nem somar rotas como prova de equivalência. Os [contratos de canais](https://github.com/wppconnect-team/wppconnect-server-go/blob/main/docs/newsletters.pt-BR.md) mostram os payloads nativos e a interpretação dos códigos HTTP.

| ID | Comportamento implementado | Limite da evidência |
| --- | --- | --- |
| Z014 | Consultar metadados por convite resolvido no WhatsApp | Teste HTTP verifica a identidade do canal; não é equivalência de todos os campos. |
| Z017, Z019 | Seguir/deixar a partir do convite | O SDK recebe o JID resolvido; a resposta contém os metadados anteriores à alteração. |
| Z022 | Enviar/remover reação por ID de servidor | 202 significa gravação no transporte; sem confirmação do cliente remoto. |
| Z023 | Registrar visualização e esperar ACK | Fixture do SDK cobre ACK, rejeição, desconexão e cancelamento; repetição com chave não duplica o efeito no teste. |
| Z025 | Consultar contadores e atualizações com `since` e `after` | Janela limitada, sem filtro `before` e sem garantia de histórico completo. |
| Z026 | Aceitar aviso específico com consentimento explícito | Testado com cliente simulado. Nenhum termo foi aceito por conta real. |
| Z027 | Enviar convite como administrador ou dono verificado | Verificação de papel e mensagem de protocolo testadas. Entrega, aceite e promoção não foram validados em duas contas. |

O [recibo de canary, pool e navegador](channels-acceptance-2026-11-10.json) não contém telefones, credenciais nem identificadores das sessões. O [plano](../PLAN.md) separa CI, publicação e aceite real. Os testes que exigem segunda conta ou Android/iPhone/Web foram adiados pelo usuário.
