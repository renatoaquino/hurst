package main

import (
	"math/rand"
	"testing"
)

func TestSplit(t *testing.T) {
	l := make([]float64, 256)
	for i := 0; i < 256; i++ {
		l[i] = rand.Float64()
	}
	j, err := Split(l, 128)
	if err != nil {
		t.Fatal(err)
	}
	if len(j) != 128 {
		t.Fatalf("Expecting 128 parts, got %d", len(j))
	}

	h := make([]float64, 11)
	for i := 0; i < 11; i++ {
		h[i] = rand.Float64()
	}
	m, err := Split(h, 2)
	if err != nil {
		t.Fatal(err)
	}
	if len(m) != 2 {
		t.Fatalf("Expecting 2 parts, got %d", len(m))
	}
	if len(m[0]) != 5 {
		t.Fatalf("Expecting 5 elements, got %d", len(m[0]))
	}
	if len(m[1]) != 6 {
		t.Fatalf("Expecting 6 elements, got %d", len(m[1]))
	}

	m, err = Split(h, 12)
	if err == nil {
		t.Fatal("Expecting a split length error, got nothing")
	}

}

func TestPartition(t *testing.T) {
	l := make([]float64, 256)
	for i := 0; i < 256; i++ {
		l[i] = rand.Float64()
	}
	h, err := Partition(l)
	if err != nil {
		t.Fatal(err)
	}
	if len(h) != 7 {
		t.Fatalf("Expecting 7 groups, got %d", len(h))
	}
	if len(h[6]) != 64 {
		t.Fatalf("Expecting 64 groups, got %d", len(h[6]))
	}
}

//From Chaos and Order in the Capital markets, Second Edition, by Edgar Peters.
//Hurst result is 0.72
var brown = []float64{
	45.47422,
	42.55601,
	46.5188,
	41.61502,
	41.77181,
	36.38034,
	39.18172,
	40.80683,
	42.62767,
	41.47926,
	35.44108,
	38.06462,
	36.23481,
	38.12881,
	38.59895,
	39.62049,
	37.94728,
	37.27003,
	35.7155,
	40.13068,
	37.78807,
	40.33214,
	38.11158,
	36.77407,
	40.23066,
	41.26233,
	41.36391,
	35.86289,
	41.85904,
	35.97116,
	38.95831,
	34.73215,
	40.23553,
	38.0261,
	37.25174,
	42.71431,
	40.50614,
	40.81272,
	41.14367,
	43.25232,
	41.56707,
	37.90438,
	40.24132,
	41.16283,
	38.63099,
	39.14979,
	38.73557,
	42.04596,
	37.13473,
	41.18287,
	35.77788,
	38.47939,
	39.83509,
	38.40837,
	36.88805,
	36.79585,
	40.23324,
	40.83193,
	41.3206,
	44.15857,
	42.32665,
	38.7613,
	38.2618,
	41.00718,
	43.57273,
	40.08365,
	41.39334,
	33.99044,
	39.99169,
	38.00431,
	40.38367,
	41.76646,
	37.84019,
	36.01189,
	36.14195,
	39.9757,
	40.00369,
	36.96267,
	40.69956,
	39.59344,
	43.26885,
	40.2914,
	41.28607,
	39.48662,
	38.98604,
	39.24064,
	42.47709,
	40.43228,
	46.32649,
	45.0697,
	41.73038,
	37.87664,
	37.60578,
	43.48858,
	35.54966,
	40.6715,
	41.70915,
	36.72274,
	43.82037,
	42.23599,
	43.99252,
	41.04976,
	39.42044,
	38.03058,
	38.23411,
	39.37609,
	38.40697,
	43.47854,
	43.39332,
	39.77525,
	40.83799,
	37.15416,
	39.52791,
	38.26106,
	40.52244,
	43.46267,
	35.96975,
	37.54511,
	42.85173,
	39.9795,
	38.79268,
	42.80823,
	45.22388,
	44.08707,
	35.95262,
	39.29638,
	38.79409,
	39.13145,
	40.65384,
	37.69769,
	38.27152,
	39.37646,
	39.64209,
	37.8354,
	39.77542,
	42.47322,
	41.42368,
	45.13499,
	45.3284,
	43.00744,
	40.96949,
	41.49458,
	39.05642,
	39.52879,
	39.8989,
	38.31439,
	37.5856,
	38.64235,
	37.6296,
	34.47122,
	33.54277,
	41.00393,
	36.86097,
	37.72265,
	39.70158,
	43.29815,
	35.819,
	34.76777,
	34.65393,
	40.79816,
	40.89328,
	35.38821,
	35.46521,
	44.98658,
	42.28198,
	38.08278,
	39.53086,
	40.35192,
	39.60774,
	36.75903,
	42.46992,
	42.9379,
	43.79816,
	36.70086,
	38.7026,
	41.68284,
	38.00106,
	38.78848,
	41.59717,
	45.17273,
	41.62175,
	39.21754,
	40.12549,
	39.24393,
	42.94091,
	35.64745,
	37.78935,
	44.39354,
	41.42976,
	37.85505,
	36.46577,
	36.66261,
	40.14956,
	41.36998,
	40.1381,
	37.39629,
	44.54419,
	42.39756,
	35.43544,
	42.64267,
	39.99272,
	43.973,
	39.37163,
	42.22032,
	40.79351,
	38.32954,
	42.0745,
	40.24001,
	40.47686,
	41.65239,
	36.90546,
	41.36013,
	38.7491,
	41.11612,
	44.79826,
	43.14703,
	38.39584,
	37.96892,
	43.34091,
	47.00554,
	41.24239,
	39.40196,
	36.67453,
	37.06655,
	38.19531,
	44.17278,
	42.31178,
	41.17061,
	47.07402,
	47.29471,
	42.72094,
	39.39516,
	38.19347,
	39.76713,
	39.60595,
	41.16065,
	36.71868,
	38.43887,
	39.55025,
	42.36142,
	43.5421,
	37.67701,
	43.06682,
	42.62402,
	36.3728,
	42.22784,
	39.21308,
	37.7949,
	40.25864,
	40.14169,
	42.25564,
	38.30389,
	35.10656,
	43.95726,
	42.58925,
	40.61366,
	41.56955,
	39.13764,
	37.28156,
	36.77036,
	31.6809,
	36.83331,
	37.14724,
	35.57775,
	39.88244,
	37.37333,
	38.01821,
	45.05355,
	39.46738,
	41.83755,
	42.38319,
	41.14247,
	45.16265,
	40.22104,
	40.36415,
	39.76302,
	42.20972,
	39.63887,
	40.74359,
	43.07138,
	40.77877,
	41.62666,
	41.95273,
	38.05709,
	37.68087,
	38.64006,
	38.4461,
	39.9015,
	42.73572,
	42.83134,
	40.33652,
	44.10318,
	38.00102,
	44.06426,
	40.4731,
	44.04457,
	40.27861,
	41.61528,
	40.53646,
	45.19104,
	38.48378,
	35.86492,
	41.2579,
	41.70177,
	38.94047,
	40.7974,
	38.16372,
	43.47877,
	42.5662,
	39.24807,
	43.36375,
	42.03311,
	41.84581,
	40.87614,
	39.55977,
	41.04644,
	40.64881,
	42.94166,
	37.37018,
	41.07071,
	38.70338,
	41.10777,
	44.54335,
	40.63602,
	38.94396,
	45.4634,
	42.38726,
	42.3946,
	39.61153,
	44.59956,
	39.32707,
	41.77308,
	37.63388,
	39.03108,
	39.46936,
	42.8321,
	36.91917,
	37.69754,
	38.74659,
	36.56134,
	38.51949,
	36.8543,
	44.6143,
	41.27391,
	37.963,
	37.07612,
	40.52002,
	37.57736,
	37.86449,
	36.85194,
	36.1528,
	36.05077,
	37.91014,
	37.17034,
	39.59316,
	33.24467,
	35.7746,
	38.21843,
	36.44605,
	43.05595,
	44.05943,
	46.80315,
	38.90282,
	39.93817,
	34.69699,
	41.37286,
	35.59125,
	36.66574,
	41.21796,
	35.49611,
	35.33717,
	39.02542,
	40.56057,
	37.88969,
	38.4017,
	35.10565,
	43.14386,
	42.76625,
	38.96064,
	40.33638,
	44.34057,
	39.65474,
	34.92081,
	40.75422,
	36.65428,
	38.792,
	44.5092,
	42.19337,
	36.94213,
	43.46756,
	41.20546,
	40.66154,
	41.94553,
	35.66089,
	35.98658,
	39.33644,
	43.26062,
	41.42673,
	42.12198,
	39.23212,
	39.27498,
	37.54181,
	42.68269,
	42.38784,
	42.87481,
	41.17289,
	39.89765,
	41.18661,
	42.68191,
	39.2564,
	41.1401,
	46.6004,
	42.02777,
	45.96446,
	44.70395,
	41.94416,
	43.73333,
	42.86214,
	44.8587,
	46.46152,
	42.13349,
	44.64348,
	40.98442,
	37.40804,
	37.48766,
	38.45201,
	41.36059,
	41.78003,
	39.9538,
	42.96399,
	40.40266,
	40.85766,
	41.55949,
	36.95011,
	42.63029,
	41.93451,
	42.45486,
	42.46915,
	41.15586,
	43.43617,
	40.14792,
	38.34502,
	42.77715,
	37.92426,
	39.73087,
	39.88557,
	41.41613,
	40.66972,
	34.62636,
	39.18851,
	39.72316,
	42.90406,
	34.46115,
	35.82703,
	38.04329,
	35.53685,
	33.78885,
	42.82296,
	37.22878,
	40.03751,
	39.00525,
	37.92589,
	43.32804,
	38.89346,
	36.88988,
	38.12963,
	41.40597,
	40.73858,
	38.12345,
	37.91511,
	40.20166,
	42.55299,
	43.92582,
	46.59028,
	42.81742,
	41.91433,
	44.37476,
	45.05684,
	43.62381,
	42.46003,
	40.93853,
	43.9466,
	45.74691,
	41.59388,
	40.00722,
	35.61822,
	35.16195,
	36.94243,
	45.26893,
	37.15262,
	40.27051,
	39.67421,
	40.03972,
	42.27792,
	45.04657,
	40.28177,
	38.67938,
	41.27745,
	39.25516,
	44.75763,
	46.37614,
	36.17558,
	44.36295,
	46.00155,
	43.44806,
	40.7477,
	37.51486,
	36.71186,
	40.78709,
	45.03045,
	43.15994,
	41.09034,
	46.20534,
	43.53859,
	41.65736,
	42.576,
	42.37403,
	42.67842,
	40.92059,
	43.0023,
	40.18782,
	41.01476,
	35.35818,
	39.25123,
	42.07744,
	37.58949,
	42.20163,
	43.22094,
	39.21471,
	41.49907,
	41.87188,
	38.47404,
	41.24654,
	40.29285,
	33.61,
	39.59441,
	42.6929,
	42.25672,
	37.45879,
	44.97766,
	39.32623,
	44.58739,
	45.01852,
	41.45086,
	40.56209,
	38.83577,
	39.61544,
	39.84228,
	39.92119,
	43.86023,
	46.85283,
	35.75996,
	40.638,
	37.13388,
	39.14868,
	37.46444,
	36.38694,
	34.67939,
	37.65355,
	37.60555,
	40.84032,
	41.67175,
	43.57566,
	36.03202,
	42.07516,
	40.90134,
	41.38979,
	41.7914,
	38.72932,
	39.29452,
	37.76977,
	38.32817,
	40.64181,
	39.52173,
	37.38647,
	38.84942,
	34.32102,
	39.49764,
	40.28845,
	42.47565,
	42.82088,
	42.277,
	37.65156,
	37.25957,
	33.80508,
	39.95499,
	39.06789,
	38.02718,
	37.56217,
	42.00036,
	37.33739,
	44.0797,
	36.81797,
	40.8698,
	40.04664,
	45.95494,
	37.71864,
	45.80341,
	40.22076,
	35.86808,
	39.6939,
	36.03485,
	44.29616,
	42.77836,
	43.67944,
	38.67348,
	42.90496,
	38.81649,
	41.69926,
	43.99026,
	47.0373,
	42.60713,
	44.96844,
	42.93016,
	42.79837,
	37.71488,
	44.80755,
	41.92877,
	41.97318,
	40.18676,
	40.75174,
	37.29416,
	41.41669,
	42.29313,
	43.76465,
	42.40368,
	40.26249,
	40.09457,
	42.90879,
	39.55097,
	35.16267,
	38.83641,
	41.69885,
	37.44271,
	41.98435,
	42.86382,
	39.75673,
	38.68184,
	45.11171,
	40.24395,
	41.09405,
	37.26426,
	39.77517,
	42.76846,
	43.00716,
	40.65384,
	37.01847,
	40.60019,
	37.54367,
	39.26697,
	39.62808,
	36.96062,
	36.25609,
	37.55869,
	41.31033,
	46.34538,
	47.08386,
	41.10557,
	46.09305,
	38.50238,
	41.23245,
	40.61206,
	39.62617,
	40.39984,
	38.37017,
	44.79132,
	38.35173,
	41.80036,
	37.76366,
	35.17949,
	37.28764,
	40.69328,
	41.56567,
	37.52876,
	36.05946,
	38.43734,
	39.66665,
	38.93267,
	38.733,
	41.45264,
	39.64677,
	40.23297,
	34.70295,
	40.26464,
	42.37408,
	42.38701,
	41.64941,
	42.96377,
	38.19841,
	39.02345,
	38.77177,
	41.51522,
	39.0166,
	38.20829,
	38.3365,
	38.75407,
	35.6497,
	40.02217,
	31.6278,
	39.1521,
	40.52456,
	40.57542,
	39.81688,
	40.70297,
	43.01102,
	42.56114,
	39.24401,
	43.54087,
	40.76357,
	40.74536,
	39.06838,
	44.83278,
	37.41712,
	45.55789,
	37.80104,
	42.48225,
	41.11916,
	44.46586,
	40.19408,
	37.6347,
	38.25857,
	40.25922,
	39.65357,
	43.49548,
	44.28566,
	35.37073,
	34.64974,
	42.31076,
	41.67571,
	40.54345,
	40.51927,
	36.22396,
	40.54132,
	39.39293,
	40.34037,
	42.45863,
	43.15921,
	36.20314,
	35.64579,
	38.34603,
	41.78037,
	35.79216,
	42.49715,
	37.64576,
	41.42941,
	36.67624,
	41.27629,
	38.2379,
	34.62199,
	41.14711,
	37.89428,
	37.45374,
	34.03383,
	38.64872,
	39.57,
	36.86844,
	43.91148,
	39.15353,
	37.45592,
	38.51285,
	40.5374,
	37.39015,
	39.61686,
	35.12693,
	34.17452,
	43.10904,
	36.98632,
	38.61848,
	37.64194,
	33.84932,
	34.94975,
	38.72701,
	37.43396,
	35.85235,
	37.75533,
	35.48315,
	36.68451,
	40.43679,
	42.07057,
	40.60281,
	37.78144,
	32.19561,
	36.57006,
	35.37629,
	38.54081,
	40.11047,
	40.38615,
	37.45115,
	40.63576,
	43.44107,
	42.8598,
	42.13489,
	37.8112,
	41.8079,
	38.93332,
	36.93653,
	39.57596,
	39.71827,
	37.87533,
	36.90723,
	43.98845,
	34.66349,
	34.95469,
	39.23472,
	37.21367,
	35.99474,
	40.22735,
	41.41341,
	39.21693,
	34.82492,
	34.51507,
	34.88574,
	40.36755,
	36.24858,
	36.02974,
	38.03783,
	39.12047,
	34.99572,
	35.759,
	41.86032,
	39.30042,
	35.27621,
	35.1936,
	36.58176,
	35.39417,
	39.48376,
	37.09365,
	39.06668,
	35.96575,
	36.43196,
	33.06973,
	39.642,
	41.26027,
	34.64079,
	34.00286,
	38.91339,
	39.71387,
	41.76277,
	39.73183,
	34.69409,
	38.03069,
	38.46732,
	35.63578,
	33.38538,
	35.63041,
	35.21228,
	34.34409,
	39.99,
	35.96511,
	33.49033,
	36.62682,
	38.22816,
	34.51382,
	33.54272,
	41.57588,
	44.12708,
	36.3208,
	37.25571,
	41.67324,
	39.43711,
	36.05919,
	42.01342,
	35.57406,
	35.05917,
	42.16963,
	39.39347,
	40.54697,
	36.39024,
	35.12942,
	37.46833,
	38.77531,
	39.74597,
	43.82986,
	42.98864,
	41.32111,
	45.74292,
	48.38778,
	45.82219,
	43.42376,
	46.39598,
	48.99336,
	42.04372,
	46.11504,
	40.65633,
	38.57129,
	38.47172,
	38.39007,
	39.88072,
	36.40014,
	43.40987,
	41.57848,
	43.9813,
	39.68525,
	34.94103,
	43.93784,
	37.45881,
	40.44743,
	40.10152,
	38.86032,
	36.14857,
	40.67007,
	38.78054,
	30.98031,
	31.29429,
	39.71067,
	41.54748,
	37.82466,
	41.92173,
	39.34708,
	40.95718,
	36.96975,
	37.86274,
	41.21288,
	42.26172,
	40.86906,
	44.51965,
	42.25937,
	40.8198,
	41.70642,
	41.77056,
	40.56536,
	35.88967,
	42.6879,
	40.84724,
	39.63992,
	41.34714,
	45.20977,
	35.23083,
	41.96883,
	43.02622,
	36.31796,
	37.83449,
	35.56613,
	37.35966,
	39.26919,
	40.55691,
	37.2283,
	41.67709,
	43.33581,
	37.54597,
	41.81398,
	41.28937,
	41.06855,
	40.08449,
	40.82016,
	41.82574,
	43.45096,
	39.51457,
	41.80509,
	40.96186,
	36.4071,
	37.99202,
	45.58952,
	36.57411,
	41.21162,
	40.21099,
	39.14748,
	39.87242,
	36.78151,
	41.21544,
	36.90537,
	39.98908,
	40.13422,
	39.88809,
	37.51876,
	44.31585,
	39.49685,
	37.28008,
	41.34757,
	38.08989,
	40.09002,
	37.82343,
	37.8083,
	41.25436,
	40.61108,
	38.61395,
	36.45138,
	37.39876,
	36.8337,
	42.45029,
	38.28586,
	40.45945,
	35.12585,
	35.3842,
	37.47065,
	36.29332,
	34.0495,
	37.47431,
	38.69455,
	35.44984,
	40.67052,
	34.07208,
	31.94135,
	39.27731,
	40.77547,
	47.02781,
	38.02953,
	40.12692,
	42.47416,
	42.21365,
	38.91418,
	40.92814,
	36.61402,
	38.29824,
	39.65691,
	39.49873,
	42.34889,
	37.93335,
	34.02641,
	40.5562,
	39.80465,
	42.24501,
	41.38593,
	40.26997,
	40.10266,
	36.47934,
	40.13805,
	40.22551,
	35.94951,
	44.90971,
	42.43009,
	42.78297,
	44.34307,
	40.70655,
}
