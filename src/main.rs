extern crate telegram_bot;

use telegram_bot::{Api, MessageType, ListeningMethod, ListeningAction};
use std::io::Read;

fn main() {
    let api = Api::from_env("TELEGRAM_TOKEN").unwrap();
    println!("getMe: {:?}", api.get_me());
    let mut listener = api.listener(ListeningMethod::LongPoll(None));

    let res = listener.listen(|u| if let Some(m) = u.message {
        let name = m.from.first_name;
        match m.msg {
            MessageType::Text(t) => {
                println!("<{}> {}", name, t);
            }
            _ => {}
        }
    });
}
