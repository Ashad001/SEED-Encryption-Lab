from Crypto.Cipher import AES
from Crypto.Util.Padding import pad
import os

# Function to encrypt data using AES-128-CBC
def aes_encrypt(plaintext, key, iv):
    cipher = AES.new(key, AES.MODE_CBC, iv)
    ciphertext = cipher.encrypt(pad(plaintext.encode('utf-8'), AES.block_size))
    return ciphertext.hex()

plaintext = "is there a bad time for pudding, is there?"
key = b'molecules#######'  # 16-byte key
iv1 = os.urandom(16)       # Generate a random 16-byte IV
iv2 = os.urandom(16)       # Generate a second random 16-byte IV

# Encrypt using two different IVs
ciphertext1 = aes_encrypt(plaintext, key, iv1)
ciphertext2 = aes_encrypt(plaintext, key, iv2)

# Encrypt using the same IV
ciphertext3 = aes_encrypt(plaintext, key, iv1)
ciphertext4 = aes_encrypt(plaintext, key, iv1)

assert ciphertext1 != ciphertext2
assert ciphertext3 == ciphertext4

print(f"Ciphertext with IV1: {ciphertext1}")
print(f"Ciphertext with IV2: {ciphertext2}")
print(f"Ciphertext with IV1 again: {ciphertext3}")
print(f"Ciphertext with IV1 again (duplicate): {ciphertext4}")
