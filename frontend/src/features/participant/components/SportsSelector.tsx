import { useState } from 'react';
import { Combobox, ComboboxInput, ComboboxOptions, ComboboxOption, ComboboxButton } from '@headlessui/react';
import { XMarkIcon, ChevronUpDownIcon, CheckIcon } from '@heroicons/react/20/solid';

type SportsSelectorProps = {
  selectedSports: string[];
  onChange: (sports: string[]) => void;
  maxSelections?: number;
};

const availableSports = [
  "サッカー",
  "野球",
  "バスケットボール", 
  "テニス",
  "バドミントン",
  "卓球",
  "バレーボール",
  "陸上競技",
  "水泳",
  "ゴルフ",
  "ラグビー",
  "アメリカンフットボール",
  "ハンドボール",
  "柔道",
  "剣道",
  "空手",
  "ボクシング",
  "レスリング",
  "体操",
  "フィギュアスケート",
];

export default function SportsSelector({ selectedSports, onChange, maxSelections = 5 }: SportsSelectorProps) {
  const [query, setQuery] = useState('');

  const filteredSports = query === ''
    ? availableSports.filter(sport => !selectedSports.includes(sport))
    : availableSports.filter(sport => 
        sport.toLowerCase().includes(query.toLowerCase()) && 
        !selectedSports.includes(sport)
      );

  const handleSelect = (sport: string) => {
    if (selectedSports.length < maxSelections && !selectedSports.includes(sport)) {
      onChange([...selectedSports, sport]);
      setQuery('');
    }
  };

  const handleRemove = (sport: string) => {
    onChange(selectedSports.filter(s => s !== sport));
  };

  return (
    <div className="space-y-3">
      {/* 選択されたスポーツ表示 */}
      {selectedSports.length > 0 && (
        <div className="flex flex-wrap gap-2">
          {selectedSports.map(sport => (
            <span
              key={sport}
              className="inline-flex items-center px-3 py-1 rounded-full text-sm bg-gray-900 text-white"
            >
              {sport}
              <button
                type="button"
                onClick={() => handleRemove(sport)}
                className="ml-2 hover:bg-gray-700 rounded-full p-0.5"
              >
                <XMarkIcon className="w-3 h-3" />
              </button>
            </span>
          ))}
        </div>
      )}

      {/* Combobox */}
      <Combobox value="" onChange={handleSelect} disabled={selectedSports.length >= maxSelections}>
        <div className="relative">
          <div className="relative w-full">
            <ComboboxInput
              className="w-full px-3 py-2 pl-3 pr-10 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-gray-500 focus:border-transparent disabled:bg-gray-100 disabled:cursor-not-allowed"
              displayValue={() => ''}
              onChange={(event) => setQuery(event.target.value)}
              placeholder={selectedSports.length >= maxSelections ? "上限に達しました" : "スポーツを検索して追加..."}
            />
            <ComboboxButton className="absolute inset-y-0 right-0 flex items-center pr-2">
              <ChevronUpDownIcon
                className="h-5 w-5 text-gray-400"
                aria-hidden="true"
              />
            </ComboboxButton>
          </div>
          <ComboboxOptions className="absolute z-10 mt-1 max-h-60 w-full overflow-auto rounded-md bg-white py-1 text-base shadow-lg ring-1 ring-black ring-opacity-5 focus:outline-none sm:text-sm">
            {filteredSports.length === 0 && query !== '' ? (
              <div className="relative cursor-default select-none py-2 px-4 text-gray-700">
                見つかりませんでした
              </div>
            ) : (
              filteredSports.map(sport => (
                <ComboboxOption
                  key={sport}
                  className={({ focus }) =>
                    `relative cursor-default select-none py-2 pl-10 pr-4 ${
                      focus ? 'bg-gray-100 text-gray-900' : 'text-gray-900'
                    }`
                  }
                  value={sport}
                >
                  {({ selected, focus }) => (
                    <>
                      <span className={`block truncate ${selected ? 'font-medium' : 'font-normal'}`}>
                        {sport}
                      </span>
                      {selected && (
                        <span
                          className={`absolute inset-y-0 left-0 flex items-center pl-3 ${
                            focus ? 'text-gray-600' : 'text-gray-600'
                          }`}
                        >
                          <CheckIcon className="h-5 w-5" aria-hidden="true" />
                        </span>
                      )}
                    </>
                  )}
                </ComboboxOption>
              ))
            )}
          </ComboboxOptions>
        </div>
      </Combobox>

      {/* 制限表示 */}
      <p className="text-xs text-gray-500">
        {selectedSports.length}/{maxSelections}件選択中
        {selectedSports.length >= maxSelections && (
          <span className="text-red-500 ml-1">（上限に達しました）</span>
        )}
      </p>
    </div>
  );
}
