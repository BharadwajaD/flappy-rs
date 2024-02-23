pub mod bird;
pub mod pipe;
pub mod shape;

use crossterm::event::{self, Event, KeyCode};

use self::{bird::Bird, pipe::Pipe};
use crate::screen::Screen;
use std::{
    io,
    time::{Duration, Instant},
};

pub trait RenderAble {
    ///On to the screen
    fn render(&self, screen: &mut dyn Screen);
}

pub struct Game<T: Screen> {
    pub screen: T,
    pub bird: Bird,
    pub pipe: Pipe,
    pub game_stats: GameStats,
}

impl<T: Screen> Game<T> {
    pub fn new(screen: T) -> Self {
        let (fxsize, fysize) = screen.get_fsize();
        return Self {
            screen,
            bird: Bird::new(0, 0),
            pipe: Pipe::new(0, 6),
            game_stats: GameStats::new(fxsize, fysize),
        };
    }

    pub fn start(&mut self) -> io::Result<()> {
        //game loop
        //self.bird.render(&mut self.screen);
        let tick_rate = Duration::from_millis(50); //get from env file
        let mut last_tick = Instant::now();
        self.pipe.render(&mut self.screen);
        let game_stats = &mut self.game_stats;

        let mut posx = 0;
        let mut posy = game_stats.fysize / 2;

        loop {
            let bird = Bird::new(posx, posy);
            let pipe = Pipe::new(posx + 10, 4);

            pipe.render(&mut self.screen);
            bird.render(&mut self.screen);

            let timeout = tick_rate.saturating_sub(last_tick.elapsed());
            if event::poll(timeout)? {
                if let Event::Key(key) = event::read()? {
                    match key.code {
                        KeyCode::Char('q') => break,
                        KeyCode::Up | KeyCode::Char('k') => {
                            game_stats.vely = 0.1;
                        }
                        _ => {}
                    }

                    if key.kind == event::KeyEventKind::Release {
                        game_stats.vely = -0.1;
                    }
                }
            }

            if last_tick.elapsed() >= tick_rate {
                last_tick = Instant::now();
            }

            game_stats.tplayed += 0.5;
            game_stats.pts += 1;
            posx = (game_stats.velx * game_stats.tplayed) as u16 % game_stats.fdiv;
            posy = (game_stats.vely * game_stats.tplayed) as u16 % game_stats.fysize;
        }

        return Ok(());
    }
}

pub struct GameStats {
    velx: f64, //increases as score increases
    vely: f64,
    pts: u16,
    fxsize: u16,
    fysize: u16,
    fdiv: u16,
    tplayed: f64,
}

impl GameStats {
    pub fn new(fxsize: u16, fysize: u16) -> Self {
        let start_velx = 0.1;
        let start_vely = -0.1;

        return Self {
            velx: start_velx,
            vely: start_vely,
            pts: 0,
            fxsize,
            fysize,
            fdiv: (fxsize as f32 * 0.8) as u16,
            tplayed: 0.0,
        };
    }
}

//TODO
pub trait IGameObject {}
