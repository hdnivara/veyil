# Veyil

Weather data in terminal.

Veyil reports brief weather data in nicely formatted line for a given ZIP code.
The output can be used in status line (vim/tmux) and/or shell prompts.


## Usage

```
$ veyil -h
veyil -- weather data in terminal
Usage: veyil [OPTIONS]
  -f FILE
        write weather data to FILE (default "/tmp/veyil.out")
  -i INTERVAL
        fetch weather data every INTERVAL seconds
  -k API-KEY
        OpenWeatherMap API-KEY
  -z ZIP-CODE
        ZIP-CODE to fetch weather data for

```

If ZIP code or API key arguments are not passed, Veyil looks for them in user's
environment using `VEYIL_ZIP_CODE` and `VEYIL_API_KEY` variables.

