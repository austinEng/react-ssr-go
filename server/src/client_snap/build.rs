extern crate cbindgen;
extern crate js_snap;

use std::process::Command;

fn main() {
  let crate_dir = std::env::var("CARGO_MANIFEST_DIR").unwrap();
  let out_dir = std::env::var("OUT_DIR").unwrap();

  Command::new("npx")
    .current_dir("../../../client")
    .args(&["webpack"])
    .status()
    .expect("webpack failed");

  Command::new("npx")
    .current_dir("../../../client/ssr")
    .args(&[
      "webpack",
      "-o", &[&out_dir, "/ssr.js"].concat()])
    .status()
    .expect("webpack failed");

  js_snap::create_snapshot(
    [&out_dir, "/ssr.js"].concat(),
    [&out_dir, "/ssr.bin"].concat());

  cbindgen::Builder::new()
    .with_crate(&crate_dir)
    .with_language(cbindgen::Language::C)
    .with_include_guard("JS_CLIENT_SNAP_H_")
    .with_item_prefix("JSClientSnap")
    .with_parse_deps(true)
    .with_parse_include(&["js_snap"])
    .generate()
    .expect("Unable to generate C bindings")
    .write_to_file(&[&crate_dir, "/gen/client_snap.h"].concat());
}