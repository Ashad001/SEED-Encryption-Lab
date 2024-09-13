#include <stdio.h>
#include <string.h>
#include <openssl/evp.h>
#include <openssl/aes.h>

// Following code is made with help of https://docs.openssl.org/1.1.1/man3/EVP_EncryptInit/#examples and CursorAI

#define KEY_SIZE 16 // AES-128 requires 16 bytes key
#define IV_SIZE 16  // AES-128-CBC uses a 16-byte IV (Initail Vector)

// Helper function to pad the key with '#' to 16 characters
void pad_key(char *key, unsigned char *padded_key)
{
    int key_len = strlen(key);
    memset(padded_key, '#', KEY_SIZE);
    memcpy(padded_key, key, key_len);
}

// Function to perform AES-128-CBC decryption
int aes_decrypt(unsigned char *ciphertext, int ciphertext_len, unsigned char *key, unsigned char *iv, unsigned char *plaintext)
{
    EVP_CIPHER_CTX *ctx;
    int len;
    int plaintext_len;

    // Create and initialize the context
    if (!(ctx = EVP_CIPHER_CTX_new()))
        return -1;

    // Initialize decryption operation with AES-128-CBC
    if (1 != EVP_DecryptInit_ex(ctx, EVP_aes_128_cbc(), NULL, key, iv))
        return -1;

    // Decrypt the ciphertext
    if (1 != EVP_DecryptUpdate(ctx, plaintext, &len, ciphertext, ciphertext_len))
        return -1;
    plaintext_len = len;

    // Finalize the decryption
    if (1 != EVP_DecryptFinal_ex(ctx, plaintext + len, &len))
        return -1;
    plaintext_len += len;

    // Clean up
    EVP_CIPHER_CTX_free(ctx);

    return plaintext_len;
}

int main()
{
    // Given data
    unsigned char ciphertext[] = {0x76, 0x4a, 0xa2, 0x6b, 0x55, 0xa4, 0xda, 0x65, 0x4d, 0xf6, 0xb1, 0x9e, 0x4b, 0xce, 0x00, 0xf4,
                                  0xed, 0x05, 0xe0, 0x93, 0x46, 0xfb, 0x0e, 0x76, 0x25, 0x83, 0xcb, 0x7d, 0xa2, 0xac, 0x93, 0xa2};
    unsigned char iv[] = {0xaa, 0xbb, 0xcc, 0xdd, 0xee, 0xff, 0x00, 0x99, 0x88, 0x77, 0x66, 0x55, 0x44, 0x33, 0x22, 0x11};
    char plaintext[] = "This is a top secret.";

    unsigned char decryptedtext[128];
    int decryptedtext_len;

    // Open the wordlist file
    FILE *wordlist = fopen("words.txt", "r");
    if (wordlist == NULL)
    {
        perror("Error opening wordlist");
        return 1;
    }

    char word[KEY_SIZE];
    unsigned char key[KEY_SIZE];

    // Loop through each word in the wordlist
    while (fgets(word, sizeof(word), wordlist))
    {
        // Remove newline character from word
        word[strcspn(word, "\n")] = 0;

        // Pad the word to 16 characters using '#'
        pad_key(word, key);

        // Decrypt the ciphertext
        decryptedtext_len = aes_decrypt(ciphertext, sizeof(ciphertext), key, iv, decryptedtext);
        if (decryptedtext_len < 0)
        {
            continue; // Decryption failed, try next key
        }

        // Null-terminate the decrypted string
        decryptedtext[decryptedtext_len] = '\0';

        // Compare decrypted text with the plaintext
        if (strcmp((char *)decryptedtext, plaintext) == 0)
        {
            printf("Key found: %s\n", word);
            break;
        }
    }

    fclose(wordlist);
    return 0;
}
