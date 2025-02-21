---
title: "Running Deepseek-R1 on Windows with Ollama"
date: 2025-02-21
description: "A comprehensive guide to installing Ollama on Windows, running the deepseek-r1:7b model, and configuring network access"
tags: ["ollama", "windows", "ai", "deepseek", "tutorial"]
---

If you’re looking for a light, straightforward way to explore large language models on Windows, Ollama is a great place to start. Rather than wrestling with dual-boot Linux or WSL drivers, you can take advantage of Ollama’s native Windows support—it’s easier to set up and still makes the most of your consumer-grade GPU.

On a system powered by an AMD Radeon RX 7800 XT, for instance, Ollama can tap into the GPU’s horsepower to speed up model inference—often beating what you’d see under WSL. In this guide, we’ll show you how to install Ollama on Windows, run the deepseek-r1:7b model, and enable local network access so you can quickly share your AI experiments with others.

### System Requirements

While Ollama can run on plenty of Windows machines, here’s one example that shows off its capabilities:

```
OS: Windows 11 Pro (Build 26100)
CPU: AMD64 Family 25 Model 97 (~4701 MHz)
RAM: 32GB Physical Memory
GPU: AMD Radeon RX 7800 XT (4GB VRAM)
```

The actual requirements may be lower, but having a dedicated GPU and sufficient RAM will significantly improve performance when running large language models.

## Installing Ollama on Windows

Installation is straightforward using Windows Package Manager (winget):

```bash
winget install Ollama.Ollama
```

## Running the deepseek-r1:7b Model

Once Ollama is installed, you can download and run the deepseek-r1:7b model with a simple command:

```bash
ollama run deepseek-r1:7b
```

### Useful Ollama Commands

Here are some helpful commands for managing your Ollama installation:

```bash
# Start the Ollama server
ollama serve

# List available models
ollama list

# If you need to force-quit Ollama
taskkill /F /IM ollama.exe
```

## Testing Local Access

You can test your Ollama installation locally using either WSL Ubuntu or your browser. Here's a simple curl command to verify it's working:

```bash
curl http://localhost:11434/api/tags
```

## Enabling Network Access

To allow other devices on your network to access your Ollama instance, you'll need to:

1. Download and install ncat from the [Nmap website](https://nmap.org/download.html)
2. Run the following command to expose Ollama to your local network:

```bash
ncat -lk 0.0.0.0 11434 --sh-exec "ncat 127.0.0.1 11434"
```

## Testing Remote Access

To test remote access, other devices on your network can use this curl command (replace IP address with your Windows machine's IP):

```bash
curl -X POST http://192.168.4.175:11434/api/generate -H "Content-Type: application/json" -d '{
  "model": "deepseek-r1:7b",
  "prompt": "Hello!",
  "stream": false
}'
```

A successful response will look something like this:

```json
{
  "model": "deepseek-r1:7b",
  "created_at": "2025-02-21T15:52:58.4296691Z",
  "response": "<think>\n\n</think>\n\nHello! How can I assist you today? 😊",
  "done": true,
  "done_reason": "stop",
  "context": [151644,9707,0,151645,151648,271,151649,271,9707,0,2585,646,358,7789,498,3351,30,26525,232],
  "total_duration": 221054700,
  "load_duration": 9989100,
  "prompt_eval_count": 5,
  "prompt_eval_duration": 2000000,
  "eval_count": 16,
  "eval_duration": 207000000
}
```

## Security Considerations

When exposing Ollama to your local network:
- Only do this on trusted networks
- Be aware that any device on your network can send requests to your Ollama instance
- Consider implementing additional security measures like a reverse proxy with authentication if needed

## Troubleshooting

If you encounter issues:
1. Ensure Ollama is running (`ollama serve`)
2. Check if the port 11434 is not being used by another application
3. Verify your Windows firewall settings allow traffic on port 11434
4. Make sure ncat is properly installed and in your system PATH
