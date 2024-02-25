pub mod bird;
pub mod consts;
pub mod pipe;
pub mod shape;

use crossterm::event::{self, Event, KeyCode};

use self::{bird::Bird, pipe::Pipe};
use crate::screen::Screen;
use std::{
    io,
    time::{Duration, Instant}, thread::sleep,
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

        let tick_rate = Duration::from_millis(consts::TICK_RATE);
        let mut last_tick = Instant::now();

        let game_stats = &mut self.game_stats;

        let mut posx = 0;
        let mut posy = game_stats.fysize / 2;

        loop {
            let bird = Bird::new(posx, posy);
            let pipe = Pipe::new(posx + consts::BIRD_PIPE_DIST, 4);

            pipe.render(&mut self.screen);
            bird.render(&mut self.screen);

            let timeout = tick_rate.saturating_sub(last_tick.elapsed());
            game_stats.vely =  consts::START_VY;
            if event::poll(timeout)? {
                if let Event::Key(key) = event::read()? {
                    match key.code {
                        KeyCode::Char('q') => break,
                        KeyCode::Up | KeyCode::Char('k') => {
                            game_stats.vely = -2.0 * consts::START_VY;
                        }
                        _ => {}
                    }

                }
            }

            if last_tick.elapsed() >= tick_rate {
                last_tick = Instant::now();
            }

            game_stats.tplayed += consts::PT_TIME_INC;
            game_stats.pts += consts::PT_SCORE_INC;

            posx = (posx as f64 +  game_stats.velx * consts::PT_TIME_INC) as u16 % game_stats.fdiv;
            posy = (posy as f64 + game_stats.vely * consts::PT_TIME_INC) as u16 % game_stats.fysize;
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
        return Self {
            velx: consts::START_VX,
            vely: consts::START_VY,
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
