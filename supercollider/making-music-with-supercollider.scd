s.boot;

// Transcribed from this talk: https://www.youtube.com/watch?v=vzspylL9Bh0

// { Pulse.ar(freq: 440, mul: 0.1, width: 0.05) }.play;
// Env.perc(releaseTime: 0.2).test.plot;


SynthDef(\blip, { |freq|
	var env = Env.perc(level: 0.1, releaseTime: 0.2).kr(doneAction: 2);
    var sig = Pulse.ar(freq: freq, width: 0.05) * env;
    Out.ar(0, [sig, sig]);
}).add;

Synth(\blip, [\freq, 440]);
Synth(\blip, [\freq, 660]);


Pdef(\melody,
    Pbind(
        \instrument, \blip,
        \dur, Pseq([0.5, 1, 0.5,  0.5,  0.5,  0.5,  0.5], inf),
        \degree, Pseq([0, 4, 6, 8, 4, 6, 7], inf),
        \scale, Scale.minor,
        \octave, 3,
    )
).play(TempoClock.default, quant: 4);

Pdef(\harmony,
    Pbind(
        \instrument, \blip,
        \dur, Pseq([0.5], inf),
        \degree, Pseq([3, 2, 1, 0], inf),
        \scale, Scale.minor,
        \octave, 6
    )
).play(TempoClock.default, quant: 4);


SynthDef(\drums, { |freq|
    var env = Env.perc(level: 0.1, releaseTime: 0.2).kr(doneAction: 2);
    var sig = WhiteNoise.ar * env;
    var trigger = Impulse.ar(freq);
    sig = Latch.ar(sig, trigger);
    Out.ar(0, [sig, sig]);
}).add;

Synth(\drums, [\freq, 500])
~bass = 500; ~snare = 10000;

Pdef(\drumBeat,
    Pbind(
        \instrument, \drums,
        \dur, Pseq([1, 1.5, 0.5, 0.5, 0.5], inf),
        \freq, Pseq([~bass, ~snare, ~bass, ~snare, ~bass], inf)
    )
).play(TempoClock.default, quant: 4);


Pdef(\harmony,
    Pbind(
        \instrument, \blip,
        \dur, Pseq([0.25, 0.25], inf),
        \degree, Pseq([3, Prand([4, 2]), 1, 0], inf),
        \scale, Scale.minor,
        \octave, 6
    )
).play(TempoClock.default, quant: 4);

