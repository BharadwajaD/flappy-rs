use std::io::{self, stdout};
use crossterm::{terminal::{enable_raw_mode, EnterAlternateScreen, disable_raw_mode, LeaveAlternateScreen}, ExecutableCommand};

use flappy_bird::{screen::CrossTermDisplay, game::Game};
use ratatui::{Terminal, backend::CrosstermBackend};

fn main() -> io::Result<()>{

    simple_logging::log_to_file("logs.txt", log::LevelFilter::Debug).unwrap();

    enable_raw_mode()?;
    stdout().execute(EnterAlternateScreen)?;

    let terminal = Terminal::new(CrosstermBackend::new(stdout()))?;
    let display = CrossTermDisplay::new(terminal);
    let mut game = Game::new(display);

    game.start()?;

    disable_raw_mode()?;
    stdout().execute(LeaveAlternateScreen)?;

    return Ok(());
}
