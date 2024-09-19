# SEED Encryption Lab

This repository contains solutions to the **SEED Encryption/Cryptography Lab** assignment, focusing on encryption concepts, common mistakes, and cryptanalysis.

## Overview
This lab covers key concepts in secret-key encryption, such as AES encryption modes, frequency analysis, padding, and common mistakes involving IV (initial vector) reuse. The tasks demonstrate how encryption vulnerabilities can be exploited and provide hands-on experience with writing crypto-related programs.

## Tasks Included:
1. **Frequency Analysis**: Breaking a monoalphabetic substitution cipher using frequency analysis.
2. **Encryption with Different Ciphers and Modes**: Testing AES-128 in various encryption modes, including CBC and CFB, with different cipher algorithms.
3. **ECB vs. CBC Modes**: Encrypting an image with ECB and CBC to observe visual differences and vulnerabilities.
4. **Padding Analysis**: Investigating PKCS#5 padding in ECB, CBC, and other encryption modes.
5. **Error Propagation in AES**: Understanding how corruption in ciphertext affects recoverability based on different encryption modes.
6. **Initial Vector (IV) Attacks**: 
   - **6.1. IV Experiment**: Exploring the necessity of unique IVs by encrypting the same plaintext with different IVs.
   - **6.2. Reusing the Same IV**: Demonstrating the weaknesses introduced by IV reuse in OFB and CFB modes.
   - **6.3. Predictable IV Attack**: Exploiting predictable IVs to reveal Bob’s secret message using AES-CBC encryption.

## Setup and Environment
The lab uses a Docker-based environment to simulate encryption oracles and execute chosen-plaintext attacks. The environment is easy to set up with the provided Docker commands.

Here’s the completed section with a bit more context:

### Docker Commands for Task 6.3
```bash
# Build the Docker containers as defined in docker-compose.yml
$ docker-compose build

# Start the containers in the background (detached mode)
$ docker-compose up -d
```

# Make sure to stop and remove the containers after use:
```bash
$ docker-compose down
```

This setup ensures the environment is properly built, launched, and cleaned up afterward.
