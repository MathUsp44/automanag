import requests
from requests.auth import HTTPBasicAuth


organization = 'sua-organizacao'  # Substitua pelo nome da sua organização no Azure DevOps
project = 'seu-projeto'  # Substitua pelo nome do seu projeto no Azure DevOps
work_item_id = 123  # Substitua pelo ID do Work Item que você deseja consultar
pat = 'seu-pat-aqui'  # Substitua pelo seu Personal Access Token


url = f'https://dev.azure.com/{organization}/{project}/_apis/wit/workitems/{work_item_id}?api-version=7.1-preview.3'


headers = {
    'Content-Type': 'application/json',
}

response = requests.get(url, headers=headers, auth=HTTPBasicAuth('', pat))


if response.status_code == 200:
    work_item = response.json()
    print("Work Item encontrado:")
    print(f"ID: {work_item['id']}")
    print(f"Título: {work_item['fields']['System.Title']}")
    print(f"Estado: {work_item['fields']['System.State']}")
else:
    print(f"Erro ao acessar o Work Item. Status Code: {response.status_code}")
