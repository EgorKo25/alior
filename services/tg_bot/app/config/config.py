import yaml

def load_config(path: str):
    with open(path, 'r') as file:
        return yaml.safe_load(file)

config = load_config('config/config.yaml')