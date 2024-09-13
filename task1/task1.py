import os
import random

with open("ciphertext.txt", "r") as file:
    ciphertext = file.read()

# Single character frequency analysis
single_char_freq = {}
for char in ciphertext:
    if char in single_char_freq and char != " " and char != "\n":
        single_char_freq[char] += 1
    else:
        single_char_freq[char] = 1

# Bigram frequency analysis
bigram_freq = {}
for i in range(len(ciphertext) - 1):
    if ciphertext[i] != " " and ciphertext[i+1] != " " and ciphertext[i] != "\n" and ciphertext[i+1] != "\n":
        bigram = ciphertext[i:i+2]
        if bigram in bigram_freq:
            bigram_freq[bigram] += 1
        else:
            bigram_freq[bigram] = 1

# Trigram frequency analysis
trigram_freq = {}
for i in range(len(ciphertext) - 2):
    if ciphertext[i] != " " and ciphertext[i+1] != " " and ciphertext[i+2] != " " and ciphertext[i] != "\n" and ciphertext[i+1] != "\n" and ciphertext[i+2] != "\n":
        trigram = ciphertext[i:i+3]
        if trigram in trigram_freq:
            trigram_freq[trigram] += 1
        else:
            trigram_freq[trigram] = 1

single_char_freq = sorted(single_char_freq.items(), key=lambda x: x[1], reverse=True)[:20]
bigram_freq = sorted(bigram_freq.items(), key=lambda x: x[1], reverse=True)[:20]
trigram_freq = sorted(trigram_freq.items(), key=lambda x: x[1], reverse=True)[:20]

import json
# save all in one file
with open("frequency_analysis.json", "w") as file:
    json.dump(
        {
            "single_char_freq": single_char_freq, 
            "bigram_freq": bigram_freq, 
            "trigram_freq": trigram_freq
        }, 
        file, 
        indent=4
    )

most_frequent_char = single_char_freq[0][0]

substitution_map = {
    'h': 'r',
    'l': 'w',
    't': 'h',  
    'a': 'c',
    'f': 'v',
    'w': 'z',
    'z': 'u',
    'v': 'a',
    'q': 's',
    'o': 'j',
    'j': 'q',
    'g': 'b',
    'r': 'g',
    'u': 'n',
    'd': 'y',
    'y': 't',
    'e': 'p',
    'n': 'e',
    'm': 'i',
    'x': 'o',
    'k': 'x',
    'p': 'd',
    'b': 'f',
    'i': 'l', 
    'c': 'm',
    's': 'k',
}

print(len(substitution_map))

missing_letters = []
for key, value in substitution_map.items():
    if key not in substitution_map.values():
        missing_letters.append(key)

print(missing_letters)


for key, value in substitution_map.items():
    key_count = {}
    value_count = {}
    for key, value in substitution_map.items():
        if key in key_count:
            key_count[key] += 1
        else:
            key_count[key] = 1
        if value in value_count:
            value_count[value] += 1
        else:
            value_count[value] = 1

key_count = {key: count for key, count in key_count.items() if count > 1}
value_count = {value: count for value, count in value_count.items() if count > 1}

decrypted_text = ""
for char in ciphertext:
    if char in substitution_map:
        decrypted_text += substitution_map[char]
    else:
        decrypted_text += char  


with open("decrypted_output.txt", "w") as output_file:
    output_file.write(decrypted_text)
