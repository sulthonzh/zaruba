from typing import List
import os

def get_env_file_names(location: str) -> List[str]:
    abs_location = os.path.abspath(location)
    env_file_names = [os.path.join(abs_location, f) for f in os.listdir(abs_location) if os.path.isfile(os.path.join(abs_location, f)) and (f.endswith('.env') or f.endswith('env.template'))]
    env_file_names.sort(key=lambda s: len(s))
    return env_file_names


def capitalize(txt: str) -> str:
    if len(txt) < 2:
        return txt.upper()
    return txt[0].upper() + txt[1:]


def snake(txt: str) -> str:
    return ''.join(['_' + ch.lower() if ch.isupper() else ch for ch in txt]).lstrip('_')


def add_python_indentation(text: str, level: int) -> str:
    spaces = (level * 4) * ' '
    indented_lines = [spaces + line for line in text.split('\n')]
    return '\n'.join(indented_lines)